package connectors

import (
	"github.com/projectjane/jane/models"
)

type Logging struct {
}

func (x Logging) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Logging) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Logging) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Logging) Help(connector models.Connector) (help string) {
	return
}
