package core

import (
	"log"
	"os"

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
			if !FileExists(config.AuditingFile) {
				nf, err := os.Create(config.AuditingFile)
				if err != nil {
					log.Fatal("ERROR: Cannot create Auditing File at "+config.AuditingFile, err)
				}
				nf.Close()
				config.Logger.Info("Created Auditing File " + config.AuditingFile)
			}
			auditing.Write(message, config)
		}

	}

}
