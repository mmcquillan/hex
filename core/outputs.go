package core

import (
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/outputs"
)

func Outputs(outputMsgs <-chan models.Message, config models.Config) {

	var cli = new(outputs.Cli)
	var slack = new(outputs.Slack)
	var auditing = new(outputs.Auditing)

	for {
		message := <-outputMsgs

		if config.CLI {
			cli.Write(message, config)
		}

		if config.Slack {
			slack.Write(message, config)
		}

		if config.Auditing {
			auditing.Write(message, config)
		}

	}

}
