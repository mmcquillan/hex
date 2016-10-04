package connectors

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
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
	for _, c := range connector.Commands {
		if match, tokens := parse.Match(c.Match, message.In.Text); match {
			if len(c.Outputs) == 0 {
				message.Out.Text = parse.Substitute(c.Output, tokens)
			} else {
				i := rand.Intn(len(c.Outputs))
				message.Out.Text = parse.Substitute(c.Outputs[i], tokens)
			}
			message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags+","+c.Tags)
			publishMsgs <- message
		}
	}
}

func (x Response) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Response) Help(connector models.Connector) (help []string) {
	help = make([]string, 0)
	for _, c := range connector.Commands {
		if !c.HideHelp {
			if c.Help != "" {
				help = append(help, c.Help)
			} else {
				help = append(help, strings.Replace(c.Match, "*", "", -1))
			}
		}
	}
	return help
}
