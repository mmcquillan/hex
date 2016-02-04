package core

import (
	"github.com/projectjane/jane/models"
)

func WhoAmI(message models.Message, publishMsgs chan<- models.Message) {
	message.Out.Text = "I know you as '" + message.In.User + "'."
	publishMsgs <- message
}
