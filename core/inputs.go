package core

import (
	"strconv"

	"github.com/hexbotio/hex/inputs"
	"github.com/hexbotio/hex/models"
)

func Inputs(inputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {

	if config.CLI {
		config.Logger.Info("Initializing CLI Input")
		var cli = new(inputs.Cli)
		go cli.Read(inputMsgs, config)
	}

	if config.Slack {
		config.Logger.Info("Initializing Slack Input")
		var slack = new(inputs.Slack)
		go slack.Read(inputMsgs, config)
	}

	if config.Scheduler {
		config.Logger.Info("Initializing Scheduler Input")
		var scheduler = new(inputs.Scheduler)
		go scheduler.Read(inputMsgs, rules, config)
	}

	if config.Webhook {
		config.Logger.Info("Initializing Webhook Input on Port " + strconv.Itoa(config.WebhookPort))
		var webhook = new(inputs.Webhook)
		go webhook.Read(inputMsgs, config)
	}

}
