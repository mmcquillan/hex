package core

import (
	"github.com/hexbotio/hex/commands"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
)

func Commands(message models.Message, outputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {
	for command, _ := range commands.List {
		if parse.Match(command, message.Attributes["hex.input"]) {
			commandService := commands.Make(command).(commands.Action)
			if commandService != nil {
				msg := deepcopy.Copy(message).(models.Message)
				commandService.Act(&msg, rules, config)
				outputMsgs <- msg
			}
		}
	}
}
