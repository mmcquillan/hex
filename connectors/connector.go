package connectors

import (
	"github.com/projectjane/jane/models"
)

type Connector interface {
	Listen(commandMsgs chan<- models.Message, connector models.Connector)
	Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector)
	Publish(publishMsgs <-chan models.Message, connector models.Connector)
	Help(connector models.Connector) (help string)
}
