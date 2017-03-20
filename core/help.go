package core

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/services"
	"sort"
	"strings"
)

func Help(message models.Message, outputMsgs chan<- models.Message, config *models.Config) {
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
				c := services.MakeService(connector.Type).(services.Service)
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
	if len(newhelp) > 0 {
		message.Out.Text = "Help for " + config.BotName + "..."
		message.Out.Detail = strings.Join(newhelp, "\n")
		outputMsgs <- message

	}
}
