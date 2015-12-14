package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"log"
	"strings"
)

func Commands(commandMsgs <-chan models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	log.Print("Initializing Commands")
	for {
		m := <-commandMsgs
		if m.In.Process {
			aliasCommands(&m, config)
			staticCommands(m, publishMsgs, config)
			for _, connector := range config.Connectors {
				if connector.Active {
					c := connectors.MakeConnector(connector.Type).(connectors.Connector)
					go c.Command(m, publishMsgs, connector)
				}
			}
		} else {
			publishMsgs <- m
		}
	}
}

func aliasCommands(message *models.Message, config *models.Config) {
	for _, alias := range config.Aliases {
		if message.In.Text == alias.Match {
			message.In.Text = alias.Output
		}
	}
}

func staticCommands(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane help" {
		Help(message, publishMsgs, config)
	}
}
