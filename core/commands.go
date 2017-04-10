package core

import (
	"strings"

	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"github.com/projectjane/jane/services"
)

func Commands(inputMsgs <-chan models.Message, outputMsgs chan<- models.Message, config *models.Config) {
	for {
		message := <-inputMsgs
		aliasCommands(&message, config)
		messages := splitCommands(message)
		for _, m := range messages {
			if m.In.Process {
				aliasCommands(&m, config)
				staticCommands(m, outputMsgs, config)
				for _, connector := range config.Connectors {
					if connector.Active {
						canRun := false
						if connector.Users == "" || connector.Users == "*" {
							canRun = true
						} else {
							users := strings.Split(connector.Users, ",")
							for _, u := range users {
								if u == m.In.User || u == m.In.Target {
									canRun = true
								}
							}
						}
						if canRun {
							c := services.MakeService(connector.Type).(services.Service)
							go c.Action(m, outputMsgs, connector)
						}
					}
				}
			} else {
				outputMsgs <- m
			}
		}
	}
}

func aliasCommands(message *models.Message, config *models.Config) {
	for _, alias := range config.Aliases {
		if match, tokens := parse.Match(alias.Match, message.In.Text); match {
			message.In.Text = parse.Substitute(alias.Output, tokens)
		}
	}
}

func staticCommands(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	if match, tokens := parse.Match(config.BotName+" help*", strings.ToLower(strings.TrimSpace(message.In.Text))); match {
		Help(message, tokens, publishMsgs, config)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == config.BotName+" ping" {
		Ping(message, outputMsgs)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == config.BotName+" whoami" {
		WhoAmI(message, outputMsgs)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == config.BotName+" passwd" {
		Passwd(message, outputMsgs)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == config.BotName+" version" {
		Version(message, outputMsgs, config)
	}
}

func splitCommands(message models.Message) (msgs []models.Message) {
	if strings.Contains(message.In.Text, "&&") {
		cmds := strings.Split(message.In.Text, "&&")
		for _, cmd := range cmds {
			var m = message
			m.In.Text = strings.TrimSpace(cmd)
			msgs = append(msgs, m)
		}
	} else {
		msgs = append(msgs, message)
	}
	return msgs
}
