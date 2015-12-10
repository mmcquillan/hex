package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"log"
)

func Listeners(commandMsgs chan<- models.Message, config *models.Config) {
	for _, connector := range config.Connectors {
		if connector.Active {
			log.Print("Initializing " + connector.ID + " listener (type: " + connector.Type + ")")
			c := connectors.MakeConnector(connector.Type).(connectors.Connector)
			go c.Listen(commandMsgs, connector)
		}
	}
}
