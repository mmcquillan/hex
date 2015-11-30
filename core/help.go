package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
)

func Help(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	message.Out.Text = "Help for jane..."
	for _, connector := range config.Connectors {
		if connector.Active {
			c := connectors.MakeConnector(connector.Type).(connectors.Connector)
			message.Out.Detail += c.Help(connector)
		}
	}
	publishMsgs <- message
}
