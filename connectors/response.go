package connectors

import (
	"github.com/projectjane/jane/models"
	"log"
	"math/rand"
	"strings"
)

type Response struct {
}

func (x Response) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Response) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if connector.Debug {
		log.Print("Incoming command message for " + connector.ID + " (" + connector.Type + ")")
		log.Printf("%+v", message)
	}
	if message.In.Process {
		txt := strings.ToLower(message.In.Text)
		for _, c := range connector.Commands {
			m := strings.ToLower(c.Match)
			ws := strings.HasPrefix(m, "*")
			we := strings.HasSuffix(m, "*")
			m = strings.Replace(m, "*", "", -1)
			eval := false
			if ws && ws && strings.Contains(txt, m) {
				eval = true
			} else if ws && !we && strings.HasSuffix(txt, m) {
				eval = true
			} else if !ws && we && strings.HasPrefix(txt, m) {
				eval = true
			} else if txt == m {
				eval = true
			}
			if eval {
				if len(c.Outputs) == 0 {
					message.Out.Text = strings.Replace(c.Output, "%msg%", strings.TrimSpace(strings.Replace(message.In.Text, m, "", 1)), -1)
				} else {
					i := rand.Intn(len(c.Outputs))
					message.Out.Text = strings.Replace(c.Outputs[i], "%msg%", strings.TrimSpace(strings.Replace(message.In.Text, m, "", 1)), -1)
				}
				publishMsgs <- message
			}
		}
	}
}

func (x Response) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Response) Help(connector models.Connector) (help string) {
	for _, c := range connector.Commands {
		if c.Help != "" {
			help += c.Help + "\n"
		} else {
			help += strings.Replace(c.Match, "*", "", -1) + "\n"
		}
	}
	return help
}
