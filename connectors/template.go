package connectors

import (
	"github.com/projectjane/jane/models"
)

type Template struct {
}

func (x Template) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Template) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Template) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Template) Help(connector models.Connector) (help string) {
	return
}
