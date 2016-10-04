package connectors

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"golang.org/x/crypto/ssh"
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
			message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags+","+command.Tags)
			message.Out.Text = connector.ID + " " + command.Name
			message.Out.Detail = parse.Substitute(command.Output, tokens)
			message.Out.Status = color
			publishMsgs <- message
		}
	}
}

func (x Exec) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Exec) Help(connector models.Connector) (help []string) {
	help = make([]string, 0)
	for _, command := range connector.Commands {
		if !command.HideHelp {
			if command.Help != "" {
				help = append(help, command.Help)
			} else {
				help = append(help, command.Match)
			}
		}
	}
	return help
}

func check(commandMsgs chan<- models.Message, command models.Command, connector models.Connector) {

	// command vars
	var state = command.Green
	var stateReset = true
	var counter = 1
	var interval = 1
	var sampling = 1
	var remind = 0
	if command.Interval > 0 {
		interval = command.Interval
	}
	if command.Sampling > 0 {
		sampling = command.Sampling
	}
	if command.Remind > 0 {
		remind = command.Remind
	}

	// loop commands
	for {

		// reset vars
		var color = "NONE"
		var match = false
		var newstate = ""
		var sendAlert = false

		// make the call
		out := callCmd(command.Cmd, command.Args, connector)

		// interpret results
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

		// handle state change
		if newstate != state {

			if stateReset {
				counter = 1
				stateReset = false
			}

			// sampling
			if counter == sampling {
				sendAlert = true
			}

			// change to green
			if newstate == command.Green {
				sendAlert = true
			}

		}

		// handle non-green state
		if newstate != command.Green && counter == remind && remind > 1 {
			sendAlert = true

		}

		// send message
		if sendAlert {
			var message models.Message
			message.In.ConnectorType = connector.Type
			message.In.ConnectorID = connector.ID
			message.In.Tags = parse.TagAppend(connector.Tags, command.Tags)
			message.In.Process = false
			message.Out.Text = connector.ID + " " + command.Name
			message.Out.Detail = strings.Replace(command.Output, "${STDOUT}", out, -1)
			message.Out.Status = color
			commandMsgs <- message
			state = newstate
			counter = 0
			stateReset = true
		}

		// wait
		counter += 1
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
	retryCounter := 1
	for retryCounter <= 3 {
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
		if serverconn {
			retryCounter = 999
		} else {
			if connector.Debug {
				log.Print("Cannot connect to server " + connector.Server + " (try #" + strconv.Itoa(retryCounter) + ")")
			}
			time.Sleep(time.Duration(3*retryCounter) * time.Second)
			retryCounter += 1
		}
	}
	if !serverconn {
		out = "ERROR - Cannot connect to server " + connector.Server
	}
	return out
}
