package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"sort"
	"strings"
)

func Help(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	message.Out.Text = "Help for " + config.BotName + "..."
	help := ""
	for _, alias := range config.Aliases {
		if !alias.HideHelp {
			if alias.Help != "" {
				help += alias.Help + "\n"
			} else {
				help += alias.Match + "\n"
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
				help += c.Help(connector)
			}
		}
	}
	helps := strings.Split(help, "\n")
	sort.Strings(helps)
	var lasthelp = ""
	var newhelps = []string{}
	for _, help := range helps {
		if help != lasthelp && help != "-" {
			newhelps = append(newhelps, help)
		}
		lasthelp = help
	}
	message.Out.Detail = strings.Join(newhelps, "\n")
	publishMsgs <- message
}
