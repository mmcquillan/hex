package connectors

import (
	"bytes"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"golang.org/x/crypto/ssh"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Exec2 struct {
}

func (x Exec2) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	var interval = 1
	if connector.Interval > 0 {
		interval = connector.Interval
	}
	var state = make(map[string]string)
	for _, chk := range connector.Checks {
		state[chk.Name] = chk.Green
	}
	for {
		check(&state, commandMsgs, connector)
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}

func (x Exec2) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	for _, c := range connector.Commands {
		if match, tokens := parse.Match(c.Match, message.In.Text); match {
			msg := strings.Replace(strings.Join(tokens, " "), "\"", "", -1)
			args := strings.Replace(c.Args, "%msg%", msg, -1)
			out := callCmd(c.Cmd, args, connector)
			message.Out.Text = strings.Replace(c.Output, "%stdout%", out, -1)
			publishMsgs <- message
		}
	}
}

func (x Exec2) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Exec2) Help(connector models.Connector) (help string) {
	for _, c := range connector.Commands {
		if c.Help != "" {
			help += c.Help + "\n"
		} else {
			help += c.Match + "\n"
		}
	}
	return help
}

func check(state *map[string]string, commandMsgs chan<- models.Message, connector models.Connector) {
	if connector.Debug {
		log.Print("Check Run for " + connector.ID)
	}
	for _, chk := range connector.Checks {
		if connector.Debug {
			log.Print("Starting " + connector.ID + " - " + chk.Name)
		}
		var color = "NONE"
		var match = false
		var newstate = ""
		out := callCmd(chk.Check, chk.Args, connector)
		if match, _ = parse.Match(chk.Green, out); match {
			newstate = chk.Green
			color = "SUCCESS"
		}
		if match, _ = parse.Match(chk.Yellow, out); match {
			newstate = chk.Yellow
			color = "WARN"
		}
		if match, _ = parse.Match(chk.Red, out); match {
			newstate = chk.Red
			color = "FAIL"
		}
		if newstate != (*state)[chk.Name] {
			var message models.Message
			message.Routes = connector.Routes
			message.In.Process = false
			message.Out.Text = connector.ID + " " + chk.Name
			message.Out.Detail = out
			message.Out.Status = color
			commandMsgs <- message
			(*state)[chk.Name] = newstate
		}
		if connector.Debug {
			log.Print("Completed " + connector.ID + " - " + chk.Name)
		}
	}
}

func callCmd(cmd string, args string, connector models.Connector) (out string) {
	if connector.Server != "" {
		out = callRemote(cmd, args, connector)
	} else {
		out = callLocal(cmd, args, connector)
	}
	return out
}

func callLocal(cmd string, args string, connector models.Connector) (out string) {
	//args = strings.Replace(args, "\"", "\\\"", -1)
	ca := cmd + " " + args
	if connector.Debug {
		log.Print("Executing: " + cmd + " " + args)
	}
	var o bytes.Buffer
	var e bytes.Buffer
	c := exec.Command("/bin/sh", "-c", ca)
	c.Stdout = &o
	c.Stderr = &e
	err := c.Run()
	if err != nil {
		log.Print(cmd + " " + args)
		log.Print(err)
		log.Print(e.String())
	}
	out = o.String()
	if connector.Debug {
		log.Print(out)
	}
	return out
}

func callRemote(cmd string, args string, connector models.Connector) (out string) {
	serverconn := true
	clientconn := &ssh.ClientConfig{
		User: connector.Login,
		Auth: []ssh.AuthMethod{
			ssh.Password(connector.Pass),
		},
	}
	port := "22"
	if connector.Port != "" {
		port = connector.Port
	}
	if connector.Debug {
		log.Print("Starting ssh connection for " + connector.Server + ":" + port)
	}
	client, err := ssh.Dial("tcp", connector.Server+":"+port, clientconn)
	if err != nil {
		log.Print(err)
	}
	if client == nil {
		serverconn = false
	} else {
		defer client.Close()
		session, err := client.NewSession()
		if err != nil {
			log.Print(err)
		}
		if session == nil {
			serverconn = false
		} else {
			defer session.Close()
			b, err := session.CombinedOutput(cmd + " " + args)
			if err != nil && connector.Debug {
				log.Print(err)
			}
			out = string(b[:])
			if connector.Debug {
				log.Print("Exec2 results for " + connector.Server + " " + cmd + " " + args + ": " + out)
			}
		}
	}
	if !serverconn {
		if connector.Debug {
			log.Print("Cannot connect to server " + connector.Server)
		}
		out = "ERROR - Cannot connect to server " + connector.Server
	}
	return out
}
