package commands

import (
	"github.com/hexbotio/hex/models"
)

type Ping struct {
}

func (x Ping) Act(message *models.Message, states map[string]models.State, config *models.Config) {
	message.Response = append(message.Response, "pong")
}
