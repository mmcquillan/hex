package core

import (
	"github.com/projectjane/jane/models"
)

func Ping(message models.Message, publishMsgs chan<- models.Message) {
	message.Out.Text = "pong"
	publishMsgs <- message
}
