package connectors

import (
	"bytes"
	"github.com/projectjane/jane/models"
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
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
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
}

func (x Exec) Publish(connector models.Connector, message models.Message, target string) {
	return
}
