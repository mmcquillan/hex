package core

import (
	"github.com/projectjane/jane/models"
	"strings"
)

func Commands(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane help" {
		Help(message, publishMsgs, config)
	}
}
