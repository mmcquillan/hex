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

type Exec struct {
}

func (x Exec) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	for _, command := range connector.Commands {
		if command.RunCheck {
			if connector.Debug {
				log.Print("Starting Listener for " + connector.ID + " " + command.Name)
			}
			go check(commandMsgs, command, connector)
		}
	}
}

func (x Exec) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	for _, command := range connector.Commands {
		if match, tokens := parse.Match(command.Match, message.In.Text); match {
			args := parse.Substitute(command.Args, tokens)
			tokens["STDOUT"] = callCmd(command.Cmd, args, connector)
			var color = "NONE"
			var match = false
			if match, _ = parse.Match(command.Green, tokens["STDOUT"]); match {
				color = "SUCCESS"
			}
			if match, _ = parse.Match(command.Yellow, tokens["STDOUT"]); match {
				color = "WARN"
			}
			if match, _ = parse.Match(command.Red, tokens["STDOUT"]); match {
				color = "FAIL"
			}
			message.Out.Text = connector.ID + " " + command.Name
			message.Out.Detail = parse.Substitute(command.Output, tokens)
			message.Out.Status = color
			publishMsgs <- message
		}
	}
}

func (x Exec) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Exec) Help(connector models.Connector) (help string) {
	for _, command := range connector.Commands {
		if !command.HideHelp {
			if command.Help != "" {
				help += command.Help + "\n"
			} else {
				help += command.Match + "\n"
			}
		}
	}
	return help
}

func check(commandMsgs chan<- models.Message, command models.Command, connector models.Connector) {
	var state = command.Green
	var interval = 1
	var remind = 0
	if command.Interval > 0 {
		interval = command.Interval
	}
	if command.Remind > 0 {
		remind = command.Remind
	}
	var counter = 0
	for {
		var color = "NONE"
		var match = false
		var newstate = ""
		counter += 1
		out := callCmd(command.Cmd, command.Args, connector)
		if match, _ = parse.Match(command.Green, out); match {
			newstate = command.Green
			color = "SUCCESS"
		}
		if match, _ = parse.Match(command.Yellow, out); match {
			newstate = command.Yellow
			color = "WARN"
		}
		if match, _ = parse.Match(command.Red, out); match {
			newstate = command.Red
			color = "FAIL"
		}
		if newstate != state || (newstate != command.Green && counter == remind && remind != 0) {
			var message models.Message
			message.Routes = connector.Routes
			message.In.Source = connector.ID
			message.In.Process = false
			message.Out.Text = connector.ID + " " + command.Name
			message.Out.Detail = strings.Replace(command.Output, "${STDOUT}", out, -1)
			message.Out.Status = color
			commandMsgs <- message
			state = newstate
		}
		if counter >= remind {
			counter = 0
		}
		time.Sleep(time.Duration(interval) * time.Minute)
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
				log.Print("Exec results for " + connector.Server + " " + cmd + " " + args + ": " + out)
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
