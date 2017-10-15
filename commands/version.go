package commands

import (
	"github.com/hexbotio/hex/models"
)

func Version(message *models.Message, config models.Config) {
	response := "Version: Non Standard Build"
	if config.Version != "" {
		response = "Version: " + config.Version
	}
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "version",
		Response: response,
	})
}
