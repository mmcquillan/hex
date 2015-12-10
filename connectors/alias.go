package connectors

import (
	"github.com/projectjane/jane/models"
)

type Alias struct {
}

func (x Alias) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Alias) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {

	return
}

func (x Alias) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Alias) Help(connector models.Connector) (help string) {
	return
}
