package commands

import (
	"github.com/hexbotio/hex/models"
)

type Version struct {
}

func (x Version) Act(message *models.Message, rules *map[string]models.Rule, config models.Config) {
	response := "Version: Non Standard Build"
	if config.Version != "" {
		response = "Version: " + config.Version
	}
	message.Outputs = append(message.Outputs, models.Output{
		Rule:      "version",
		StartTime: models.MessageTimestamp(),
		EndTime:   models.MessageTimestamp(),
		Response:  response,
	})
}
