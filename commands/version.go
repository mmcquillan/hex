package commands

import (
	"github.com/hexbotio/hex/models"
)

type Version struct {
}

func (x Version) Act(message *models.Message, states map[string]models.State, config *models.Config) {
	if config.Version != "" {
		message.Response = append(message.Response, "Version: "+config.Version)
	} else {
		message.Response = append(message.Response, "Version: Non Standard Build")
	}
}
