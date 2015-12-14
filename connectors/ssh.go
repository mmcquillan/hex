package connectors

import (
	"github.com/projectjane/jane/models"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

type Ssh struct {
}

func (x Ssh) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Ssh) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				msg := strings.TrimSpace(strings.Replace(message.In.Text, c.Match, "", 1))
				msg = strings.Replace(msg, "\"", "", -1)
				args := strings.Replace(c.Args, "%msg%", msg, -1)
				out := callSsh(c.Cmd, args, connector)
				message.Out.Text = strings.Replace(c.Output, "%stdout%", out, -1)
				publishMsgs <- message
			}
		}
	}
}

func (x Ssh) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Ssh) Help(connector models.Connector) (help string) {
	for _, c := range connector.Commands {
		if c.Help != "" {
			help += c.Help + "\n"
		} else {
			help += c.Match + "\n"
		}
	}
	return help
}

func callSsh(cmd string, args string, connector models.Connector) (out string) {
	serverconn := true
	clientconn := &ssh.ClientConfig{
		User: connector.Login,
		Auth: []ssh.AuthMethod{
			ssh.Password(connector.Pass),
		},
	}
	if connector.Debug {
		log.Print("Starting ssh connection for " + connector.Server)
	}
	client, err := ssh.Dial("tcp", connector.Server+":22", clientconn)
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
				log.Print("Ssh results for " + connector.Server + " " + cmd + " " + args + ": " + out)
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
