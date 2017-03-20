package services

import (
	"github.com/projectjane/jane/models"
)

type Template struct {
}

func (x Template) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Template) Command(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Template) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Template) Help(connector models.Connector) (help []string) {
	return
}
