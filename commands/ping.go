package commands

import (
	"github.com/mmcquillan/hex/models"
)

func Ping(message *models.Message) {
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "ping",
		Response: "pong",
		Success:  true,
	})
}
