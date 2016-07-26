package core

import (
	"github.com/projectjane/jane/models"
)

func Version(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	message.Out.Text = "Jane Version: " + config.Version
	publishMsgs <- message
}
