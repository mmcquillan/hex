package core

import (
	"github.com/hexbotio/hex/inputs"
	"github.com/hexbotio/hex/models"
)

func Inputs(inputMsgs chan<- models.Message, config models.Config) {

	// cli
	if config.CLI {
		startInput("cli", inputMsgs, config)
	}

	// slack
	if config.Slack {
		startInput("slack", inputMsgs, config)
	}

	// scheduler
	if config.Scheduler {
		startInput("scheduler", inputMsgs, config)
	}

	// webhook
	if config.Webhook {
		startInput("webhook", inputMsgs, config)
	}

}

func startInput(service string, inputMsgs chan<- models.Message, config models.Config) {
	if inputs.Exists(service) {
		inputService := inputs.Make(service).(inputs.Input)
		if inputService != nil {
			config.Logger.Info("Initializing Input for " + service)
			go inputService.Read(inputMsgs, config)
		}
	}
}
