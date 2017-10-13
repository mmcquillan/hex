package core

import (
	"github.com/hexbotio/hex/inputs"
	"github.com/hexbotio/hex/models"
)

func Inputs(inputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {

	if config.CLI {
		var cli = new(inputs.Cli)
		go cli.Read(inputMsgs, config)
	}

	if config.Slack {
		var slack = new(inputs.Slack)
		go slack.Read(inputMsgs, config)
	}

	if config.Scheduler {
		var scheduler = new(inputs.Scheduler)
		go scheduler.Read(inputMsgs, rules, config)
	}

	if config.Webhook {
		var webhook = new(inputs.Webhook)
		go webhook.Read(inputMsgs, config)
	}

}
