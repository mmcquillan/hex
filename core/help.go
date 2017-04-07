package core

import (
	"fmt"
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"sort"
	"strings"
)

func Help(message models.Message, tokens map[string]string, publishMsgs chan<- models.Message, config *models.Config) {

	help := make([]string, 0)

	// pull all help from the aliases
	for _, alias := range config.Aliases {
		if !alias.HideHelp {
			if alias.Help != "" {
				help = append(help, alias.Help)
			} else {
				help = append(help, alias.Match)
			}
		}
	}

	// pull all help from the connectors
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

	// sort, filter and de-dupe help
	sort.Strings(help)
	var lasthelp = ""
	var newhelp = make([]string, 0)
	for _, h := range help {
		if (tokens["1"] != "" && strings.Contains(h, tokens["1"])) || tokens["1"] == "" {
			if h != lasthelp {
				newhelp = append(newhelp, h)
			}
			lasthelp = h
		}
	}

	fmt.Printf("%+v", tokens)

	// output help
	if len(newhelp) > 0 {
		message.Out.Text = "Help for " + config.BotName + "..."
		message.Out.Detail = strings.Join(newhelp, "\n")
		publishMsgs <- message

	}
}
