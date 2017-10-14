package commands

import (
	"github.com/hexbotio/hex/models"
)

func Ping(message *models.Message) {
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "ping",
		Response: "pong",
	})
}
