package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
	"strings"
)

func Publishers(publishMsgs <-chan models.Message, config *models.Config) {
	log.Print("Initializing Publishers")
	var chans = make(map[string]chan models.Message)
	for _, connector := range config.Connectors {
		if connector.Active {
			chans[connector.ID] = make(chan models.Message)
			c := connectors.MakeConnector(connector.Type).(connectors.Connector)
			go c.Publish(chans[connector.ID], connector)
		}
	}
	for {
		message := <-publishMsgs
		for _, route := range config.Routes {
			for _, m := range route.Matches {
				match := true
				match, _ = parse.Match(m.Message, message.Out.Text+" "+message.Out.Detail)
				matchRoute(&match, message.In.ConnectorType, m.ConnectorType)
				matchRoute(&match, message.In.ConnectorID, m.ConnectorID)
				matchRouteTags(&match, message.In.Tags, m.Tags)
				matchRoute(&match, message.In.Target, m.Target)
				matchRoute(&match, message.In.User, m.User)
				if match {
					for _, connector := range config.Connectors {
						if connector.Active && sendToConnector(connector.ID, route.Connectors) {
							message.Out.Target = route.Targets
							chans[connector.ID] <- message
						}
					}
				}
			}
		}
	}
}

func matchRoute(match *bool, mValue string, rValue string) {
	if *match {
		*match = parse.SimpleMatch(mValue, rValue)
	}
}

func matchRouteTags(match *bool, mValue string, rValue string) {
	if *match {
		*match = parse.TagMatch(mValue, rValue)
	}
}

func sendToConnector(connId string, connectors string) (send bool) {
	send = false
	if connectors == "*" {
		send = true
	}
	r := strings.Split(connectors, ",")
	for _, v := range r {
		if v == connId {
			send = true
		}
	}
	return send
}
