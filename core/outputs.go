package core

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/outputs"
	"github.com/projectjane/jane/parse"
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
							target = message.Inputs["jane.input"]
						}
						message.Inputs["jane.output"] = target
						outputChannels[serviceOutput] <- message
					}
				}
			}
		}
	}
}
