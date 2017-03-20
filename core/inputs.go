package core

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/services"
	"log"
)

func Inputs(inputMsgs chan<- models.Message, config *models.Config) {
	for _, connector := range config.Connectors {
		if connector.Active {
			log.Print("Initializing " + connector.Type + " input: " + connector.ID)
			c := services.MakeService(connector.Type).(services.Service)
			go c.Input(inputMsgs, connector)
		}
	}
}
