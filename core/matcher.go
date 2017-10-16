package core

import (
	"strconv"
	"strings"

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

			// match for input
			if rule.Active && rule.Match != "" && parse.Match(rule.Match, message.Attributes["hex.input"]) {
				if parse.Member(rule.ACL, message.Attributes["hex.user"]) || parse.Member(rule.ACL, message.Attributes["hex.channel"]) {
					config.Logger.Debug("Matched Rule '" + rule.Name + "' with input '" + message.Attributes["hex.input"] + "'")
					msg := deepcopy.Copy(message).(models.Message)
					go runRule(rule, msg, outputMsgs, *plugins, config)
				}
			}

			// match for schedule
			if rule.Active && rule.Schedule != "" && rule.Schedule == message.Attributes["hex.schedule"] {
				config.Logger.Debug("Matched Rule '" + rule.Name + "' with schedule '" + message.Attributes["hex.schedule"] + "'")
				msg := deepcopy.Copy(message).(models.Message)
				go runRule(rule, msg, outputMsgs, *plugins, config)
			}

			// match for webhook
			if rule.Active && rule.URL != "" && parse.Match(rule.URL, message.Attributes["hex.url"]) {
				config.Logger.Debug("Matched Rule '" + rule.Name + "' with url '" + message.Attributes["hex.url"] + "'")
				msg := deepcopy.Copy(message).(models.Message)
				go runRule(rule, msg, outputMsgs, *plugins, config)
			}

		}
	}
}

func runRule(rule models.Rule, message models.Message, outputMsgs chan<- models.Message, plugins map[string]models.Plugin, config models.Config) {
	message.Attributes["hex.rule.runid"] = models.MessageID()
	message.Attributes["hex.rule.name"] = rule.Name
	message.Attributes["hex.rule.format"] = strconv.FormatBool(rule.Format)
	message.Attributes["hex.rule.channel"] = rule.Channel
	actionCounter := 0
	lastAction := true
	lastConfig := rule.Actions[0].Config
	for _, action := range rule.Actions {
		if lastAction || action.RunOnFail {
			if _, exists := plugins[action.Type]; exists {
				startTime := models.MessageTimestamp()
				attrName := "hex.output." + strconv.Itoa(actionCounter)
				if action.LastConfig {
					action.Config = lastConfig
				}
				for key, _ := range action.Config {
					action.Config[key] = parse.Substitute(action.Config[key], message.Attributes)
				}
				cmd := parse.Substitute(action.Command, message.Attributes)
				args := hexplugin.Arguments{
					Debug:   rule.Debug || config.Debug,
					Command: cmd,
					Config:  action.Config,
				}
				resp := plugins[action.Type].Action.Perform(args)
				if action.OutputToVar {
					message.Attributes[attrName+".response"] = strings.TrimSpace(resp.Output)
				} else {
					message.Outputs = append(message.Outputs, models.Output{
						Rule:     rule.Name,
						Response: resp.Output,
						Success:  resp.Success,
						Command:  cmd,
					})
				}
				lastAction = resp.Success
				lastConfig = action.Config
				message.Attributes[attrName+".duration"] = strconv.FormatInt(models.MessageTimestamp()-startTime, 10)
			} else {
				config.Logger.Error("Missing Plugin " + action.Type)
			}
		}
		actionCounter += 1
	}
	message.EndTime = models.MessageTimestamp()
	outputMsgs <- message
}
