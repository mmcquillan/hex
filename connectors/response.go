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

		// loop through and find command matches
		var match []string
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				out := strings.Replace(c.Output, "%msg%", strings.TrimSpace(strings.Replace(message.In.Text, c.Match, "", 1)), -1)
				match = append(match, out)
			}
		}

		// if no match, just leave
		if len(match) == 0 {
			return
		}

		// if more than one match, pick a random one
		var i = 0
		if len(match) > 0 {
			i = rand.Intn(len(match))
		}
		if connector.Debug {
			log.Print("Match: " + match[i])
		}

		message.Out.Text = match[i]
		publishMsgs <- message

	}
}

func (x Response) Publish(connector models.Connector, message models.Message, target string) {
	return
}
