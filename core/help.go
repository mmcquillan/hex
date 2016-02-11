package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"sort"
	"strings"
)

func Help(message models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	message.Out.Text = "Help for jane..."
	help := ""
	for _, alias := range config.Aliases {
		help += alias.Match + "\n"
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
	message.Out.Detail = strings.Join(helps, "\n")
	publishMsgs <- message
}
