package commands

import (
	"github.com/hexbotio/hex/models"
)

type Ping struct {
}

func (x Ping) Act(message *models.Message, config *models.Config) {
	message.Response = append(message.Response, "pong")
}
