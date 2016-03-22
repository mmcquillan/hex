package connectors

import (
	"bytes"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
	"os/exec"
	"strings"
)

type Exec struct {
}

func (x Exec) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Exec) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	for _, c := range connector.Commands {
		if match, _ := parse.Match(c.Match, message.In.Text); match {
			msg := strings.TrimSpace(strings.Replace(message.In.Text, c.Match, "", 1))
			msg = strings.Replace(msg, "\"", "", -1)
			args := strings.Split(strings.Replace(c.Args, "%msg%", msg, -1), " ")
			cmd := exec.Command(c.Cmd, args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Print(c.Cmd + " " + strings.Join(args, " "))
				log.Print(err)
			}
			message.Out.Text = strings.Replace(c.Output, "%stdout%", out.String(), -1)
			publishMsgs <- message
		}
	}
}

func (x Exec) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Exec) Help(connector models.Connector) (help string) {
	for _, c := range connector.Commands {
		if c.Help != "" {
			help += c.Help + "\n"
		} else {
			help += c.Match + "\n"
		}
	}
	return help
}
