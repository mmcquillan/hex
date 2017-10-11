package core

import (
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/outputs"
)

func Outputs(outputMsgs <-chan models.Message, config models.Config) {
	for {
		message := <-outputMsgs
		if config.CLI {
			var cli outputs.Cli
			cli.Write(message, config)
		}
		if config.Slack {
			var slack outputs.Slack
			slack.Write(message, config)
		}
		if config.Auditing {
			var auditing outputs.Auditing
			auditing.Write(message, config)
		}
	}
}

func startOutput(service string, outputMsgs <-chan models.Message, config models.Config) {
	if outputs.Exists(service) {
		outputService := outputs.Make(service).(outputs.Output)
		if outputService != nil {
			config.Logger.Info("Initializing Output for " + service)
			go outputService.Write(outputMsgs, config)
		}
	}
}
