package commands

import (
	"github.com/hexbotio/hex/models"
)

type Whoami struct {
}

func (x Whoami) Act(message *models.Message, config *models.Config) {
	message.Response = append(message.Response, message.Inputs["hex.user"])
}
