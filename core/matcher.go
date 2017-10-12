package core

import (
	"strconv"

	"github.com/hexbotio/hex-plugin"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
)

func Matcher(inputMsgs <-chan models.Message, outputMsgs chan<- models.Message, plugins *map[string]models.Plugin, rules *map[string]models.Rule, config models.Config) {
	for {
		message := <-inputMsgs
		Commands(message, outputMsgs, rules, config)
		for _, rule := range *rules {
			if rule.Active && parse.Match(rule.Match, message.Attributes["hex.input"]) {
				if parse.Member(rule.ACL, message.Attributes["hex.user"]) || parse.Member(rule.ACL, message.Attributes["hex.channel"]) {
					config.Logger.Debug("Matched Rule '" + rule.Name + "' with input '" + message.Attributes["hex.input"] + "'")
					msg := deepcopy.Copy(message).(models.Message)
					go runRule(rule, msg, outputMsgs, *plugins, config)
				}
			}
		}
	}
}

func runRule(rule models.Rule, message models.Message, outputMsgs chan<- models.Message, plugins map[string]models.Plugin, config models.Config) {
	message.Attributes["hex.rule.runid"] = models.MessageID()
	message.Attributes["hex.rule.name"] = rule.Name
	message.Attributes["hex.rule.format"] = strconv.FormatBool(rule.Format)
	for _, action := range rule.Actions {
		if _, exists := plugins[action.Type]; exists {
			args := hexplugin.Arguments{
				Debug:   rule.Debug || config.Debug,
				Command: parse.Substitute(action.Command, message.Attributes),
				Config:  action.Config,
			}
			resp := plugins[action.Type].Action.Perform(args)
			message.Outputs = append(message.Outputs, models.Output{
				Rule:      rule.Name,
				StartTime: models.MessageTimestamp(),
				EndTime:   models.MessageTimestamp(),
				Response:  resp.Output,
				Success:   resp.Success,
			})
		} else {
			config.Logger.Error("Missing Plugin " + action.Type)
		}
	}
	outputMsgs <- message
}
