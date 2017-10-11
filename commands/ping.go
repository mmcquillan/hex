package commands

import (
	"github.com/hexbotio/hex/models"
)

type Ping struct {
}

func (x Ping) Act(message *models.Message, rules *map[string]models.Rule, config models.Config) {
	message.Outputs = append(message.Outputs, models.Output{
		Rule:      "ping",
		StartTime: models.MessageTimestamp(),
		EndTime:   models.MessageTimestamp(),
		Response:  "pong",
	})
}
