package core

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hexbotio/hex-plugin"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/mohae/deepcopy"
)

// Matcher function
func Matcher(inputMsgs <-chan models.Message, outputMsgs chan<- models.Message, plugins *map[string]models.Plugin, rules *map[string]models.Rule, config models.Config) {
	state := make(map[string]bool)
	for _, rule := range *rules {
		state[rule.Id] = true
	}
	for {
		message := <-inputMsgs
		match := false
		config.Logger.Debug("Matcher - Eval of Message ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
		if parse.EitherMember(config.ACL, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
			Commands(message, outputMsgs, rules, config)
		}
		for _, rule := range *rules {

			// match for input
			if rule.Active && rule.Match != "" && parse.Match(rule.Match, message.Attributes["hex.input"]) {
				if parse.EitherMember(config.ACL, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
					if parse.EitherMember(rule.ACL, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
						match = true
						config.Logger.Debug("Matcher - Matched Rule '" + rule.Name + "' with input '" + message.Attributes["hex.input"] + "' on ID:" + message.Attributes["hex.id"])
						config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
						msg := deepcopy.Copy(message).(models.Message)
						go runRule(rule, msg, outputMsgs, state, *plugins, config)
					}
				}
			}

			// match for schedule
			if rule.Active && rule.Schedule != "" && rule.Schedule == message.Attributes["hex.schedule"] {
				match = true
				config.Logger.Debug("Matcher - Matched Rule '" + rule.Name + "' with schedule '" + message.Attributes["hex.schedule"] + "' on ID:" + message.Attributes["hex.id"])
				config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
				msg := deepcopy.Copy(message).(models.Message)
				go runRule(rule, msg, outputMsgs, state, *plugins, config)
			}

			// match for webhook
			if rule.Active && rule.URL != "" && parse.Match(rule.URL, message.Attributes["hex.url"]) {
				match = true
				config.Logger.Debug("Matcher - Matched Rule '" + rule.Name + "' with url '" + message.Attributes["hex.url"] + "' on ID:" + message.Attributes["hex.id"])
				config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
				msg := deepcopy.Copy(message).(models.Message)
				go runRule(rule, msg, outputMsgs, state, *plugins, config)
			}

		}
		if !match && message.Attributes["hex.service"] == "command" {
			StopPlugins(*plugins, config)
			os.Exit(0)

		}
	}
}

func runRule(rule models.Rule, message models.Message, outputMsgs chan<- models.Message, state map[string]bool, plugins map[string]models.Plugin, config models.Config) {
	config.Logger.Debug("Matcher - Running Rule " + rule.Name + " for ID:" + message.Attributes["hex.id"])
	config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
	message.Attributes["hex.rule.runid"] = models.MessageID()
	message.Attributes["hex.rule.name"] = rule.Name
	message.Attributes["hex.rule.format"] = strconv.FormatBool(rule.Format)
	message.Attributes["hex.rule.channel"] = rule.Channel
	for key, value := range config.Vars {
		message.Attributes["hex.var."+key] = value
	}
	actionCounter := 0
	ruleResult := true
	lastAction := true
	lastConfig := rule.Actions[0].Config
	for _, action := range rule.Actions {
		config.Logger.Debug("Matcher - Evaluating Action " + rule.Name + "." + action.Type + " [" + strconv.Itoa(actionCounter) + "] for ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
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
				if !resp.Success {
					ruleResult = false
				}
				lastAction = resp.Success
				lastConfig = action.Config
				message.Attributes[attrName+".duration"] = strconv.FormatInt(models.MessageTimestamp()-startTime, 10)
				if action.OutputToVar {
					message.Attributes[attrName+".response"] = strings.TrimSpace(resp.Output)
				} else if !action.HideOutput {
					message.Outputs = append(message.Outputs, models.Output{
						Rule:     rule.Name,
						Response: resp.Output,
						Success:  resp.Success,
						Command:  cmd,
					})
				}
			} else {
				config.Logger.Error("Matcher - Missing Plugin " + action.Type)
			}
		}
		actionCounter += 1
	}
	message.EndTime = models.MessageTimestamp()
	if !rule.OutputOnChange && (!rule.OutputFailOnly || !ruleResult) {
		config.Logger.Debug("Matcher - Output ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
		outputMsgs <- message
	} else if rule.OutputOnChange && ruleResult != state[rule.Id] {
		config.Logger.Debug("Matcher - Output Change ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
		outputMsgs <- message
	} else {
		config.Logger.Debug("Matcher - Discarding ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
	}
	state[rule.Id] = ruleResult
}
