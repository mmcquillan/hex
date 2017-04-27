package core

import (
	"github.com/projectjane/jane/inputs"
	"github.com/projectjane/jane/models"
	"log"
)

func Inputs(inputMsgs chan<- models.Message, config *models.Config) {
	for _, service := range config.Services {
		if service.Active && inputs.Exists(service.Type) {
			inputService := inputs.Make(service.Type).(inputs.Input)
			if inputService != nil {
				log.Print("Initializing Input " + service.Type + ": " + service.Name)
				go inputService.Read(inputMsgs, service)
			}
		}
	}
}
