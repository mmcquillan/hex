package services

import (
	"github.com/projectjane/jane/models"
)

// Service Interface representing a connector. Must have Input, Command, Output and Help
type Service interface {
	Input(inputMsgs chan<- models.Message, connector models.Connector)
	Action(message models.Message, outputMsgs chan<- models.Message, connector models.Connector)
	Output(outputMsgs <-chan models.Message, connector models.Connector)
	Help(connector models.Connector) (help []string)
}
