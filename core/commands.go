package core

import (
	"log"
	"strings"

	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

func Commands(commandMsgs <-chan models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	for {
		message := <-commandMsgs
		aliasCommands(&message, config)
		messages := splitCommands(message)
		for _, m := range messages {
			if m.In.Process {
				aliasCommands(&m, config)
				staticCommands(m, publishMsgs, config)
				for _, connector := range config.Connectors {
					if connector.Active {
						canRun := false
						if connector.Users == "" || connector.Users == "*" {
							canRun = true
						} else {
							users := strings.Split(connector.Users, ",")
							for _, u := range users {
								if u == m.In.User {
									canRun = true
								}
							}
						}
						if canRun {
							c := connectors.MakeConnector(connector.Type).(connectors.Connector)
							go c.Command(m, publishMsgs, connector)
						}
					}
				}
			} else {
				publishMsgs <- m
			}
		}
	}
}

func aliasCommands(message *models.Message, config *models.Config) {
	log.Printf("%+v", message)
	for _, alias := range config.Aliases {
		if match, tokens := parse.Match(alias.Match, message.In.Text); match {
			log.Printf("%+v", tokens)
			message.In.Text = parse.Substitute(alias.Output, tokens)
			log.Printf("%+v", message)
		}
	}
}

func staticCommands(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane help" {
		Help(message, publishMsgs, config)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane whoami" {
		WhoAmI(message, publishMsgs)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane passwd" {
		Passwd(message, publishMsgs)
	}
	if strings.ToLower(strings.TrimSpace(message.In.Text)) == "jane version" {
		Version(message, publishMsgs, config)
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
