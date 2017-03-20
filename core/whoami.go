package core

import (
	"github.com/projectjane/jane/models"
)

func WhoAmI(message models.Message, outputMsgs chan<- models.Message) {
	message.Out.Text = "I know you as '" + message.In.User + "'."
	outputMsgs <- message
}
