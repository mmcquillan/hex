package core

import (
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/outputs"
	"github.com/hexbotio/hex/parse"
	"log"
	"strings"
)

func Outputs(outputMsgs <-chan models.Message, config *models.Config) {
	var outputChannels = make(map[string]chan models.Message)
	for _, service := range config.Services {
		if service.Active && outputs.Exists(service.Type) {
			outputChannels[service.Name] = make(chan models.Message)
			outputService := outputs.Make(service.Type).(outputs.Output)
			if outputService != nil {
				log.Print("Initializing Output " + service.Type + ": " + service.Name)
				go outputService.Write(outputChannels[service.Name], service)

			}
		}
	}
	for {
		message := <-outputMsgs
		for _, output := range message.Outputs {
			for serviceOutput, _ := range outputChannels {
				if parse.Match(serviceOutput, output.Name) {
					for _, target := range strings.Split(output.Targets, ",") {
						if target == "*" {
							target = message.Inputs["hex.input"]
						}
						message.Inputs["hex.output"] = target
						outputChannels[serviceOutput] <- message
					}
				}
			}
		}
	}
}
