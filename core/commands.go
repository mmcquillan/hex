package core

import (
	"github.com/hexbotio/hex/commands"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
)

func Commands(message models.Message, outputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {

	if parse.Match("help*", message.Attributes["hex.input"]) {
		msg := deepcopy.Copy(message).(models.Message)
		commands.Help(&msg, rules, config)
		outputMsgs <- msg
	}

	if parse.Match("version", message.Attributes["hex.input"]) {
		msg := deepcopy.Copy(message).(models.Message)
		commands.Version(&msg, config)
		outputMsgs <- msg
	}

	if parse.Match("ping", message.Attributes["hex.input"]) {
		msg := deepcopy.Copy(message).(models.Message)
		commands.Ping(&msg)
		outputMsgs <- msg
	}

	if parse.Match("rules", message.Attributes["hex.input"]) {
		if parse.Member(config.Admins, message.Attributes["hex.user"]) || parse.Member(config.Admins, message.Attributes["hex.channel"]) {
			msg := deepcopy.Copy(message).(models.Message)
			commands.Rules(&msg, rules, config)
			outputMsgs <- msg
		}
	}

}
