package core

import (
	"os"

	"github.com/mmcquillan/hex/models"
	"github.com/mmcquillan/hex/outputs"
)

func Outputs(outputMsgs <-chan models.Message, plugins *map[string]models.Plugin, config models.Config) {

	var command = new(outputs.Command)
	var cli = new(outputs.Cli)
	var slack = new(outputs.Slack)
	var auditing = new(outputs.Auditing)

	for {
		message := <-outputMsgs

		if config.Command != "" {
			command.Write(message, config)
			StopPlugins(*plugins, config)
			for _, output := range message.Outputs {
				if !output.Success {
					os.Exit(1)
				}
			}
			os.Exit(0)
		}

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
