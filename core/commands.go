package core

import (
	"github.com/hexbotio/hex/commands"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
)

func Commands(message models.Message, outputMsgs chan<- models.Message, rules *map[string]models.Rule, config models.Config) {

	if parse.Match("help*", message.Attributes["hex.input"]) {
		config.Logger.Debug("Matcher Command - Help Match ID:" + message.Attributes["hex.id"])
		msg := deepcopy.Copy(message).(models.Message)
		commands.Help(&msg, rules, config)
		msg.EndTime = models.MessageTimestamp()
		outputMsgs <- msg
	}

	if parse.Match("version", message.Attributes["hex.input"]) {
		config.Logger.Debug("Matcher Command - Version Match ID:" + message.Attributes["hex.id"])
		msg := deepcopy.Copy(message).(models.Message)
		commands.Version(&msg, config)
		msg.EndTime = models.MessageTimestamp()
		outputMsgs <- msg
	}

	if parse.Match("ping", message.Attributes["hex.input"]) {
		config.Logger.Debug("Matcher Command - Ping Match ID:" + message.Attributes["hex.id"])
		msg := deepcopy.Copy(message).(models.Message)
		commands.Ping(&msg)
		msg.EndTime = models.MessageTimestamp()
		outputMsgs <- msg
	}

	if parse.Match("rules", message.Attributes["hex.input"]) {
		if parse.EitherMember(config.Admins, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
			config.Logger.Debug("Matcher Command - Rules Match ID:" + message.Attributes["hex.id"])
			msg := deepcopy.Copy(message).(models.Message)
			commands.Rules(&msg, rules, config)
			msg.EndTime = models.MessageTimestamp()
			outputMsgs <- msg
		}
	}

}
