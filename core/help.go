package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"sort"
	"strings"
)

func Help(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	message.Out.Text = "Help for " + config.BotName + "..."
	help := make([]string, 0)
	for _, alias := range config.Aliases {
		if !alias.HideHelp {
			if alias.Help != "" {
				help = append(help, alias.Help)
			} else {
				help = append(help, alias.Match)
			}
		}
	}
	for _, connector := range config.Connectors {
		if connector.Active {
			canRun := false
			if connector.Users == "" || connector.Users == "*" {
				canRun = true
			} else {
				users := strings.Split(connector.Users, ",")
				for _, u := range users {
					if u == message.In.User {
						canRun = true
					}
				}
			}
			if canRun {
				c := connectors.MakeConnector(connector.Type).(connectors.Connector)
				help = append(help, c.Help(connector)...)
			}
		}
	}
	sort.Strings(help)
	var lasthelp = ""
	var newhelp = make([]string, 0)
	for _, h := range help {
		if h != lasthelp {
			newhelp = append(newhelp, h)
		}
		lasthelp = h
	}
	message.Out.Detail = strings.Join(newhelp, "\n")
	publishMsgs <- message
}
