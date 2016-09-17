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
	for {
		message := <-publishMsgs
		for _, route := range config.Routes {
			for _, m := range route.Matches {
				match := true
				match, _ = parse.Match(m.Message, message.Out.Text+" "+message.Out.Detail)
				match = matchRoute(match, message.In.ConnectorType, m.ConnectorType)
				match = matchRoute(match, message.In.ConnectorID, m.ConnectorID)
				match = matchRoute(match, message.In.Target, m.Target)
				match = matchRoute(match, message.In.User, m.User)
				if match {
					for _, connector := range config.Connectors {
						if connector.Active {
							if sendToConnector(connector.ID, route.Connectors) {
								if connector.Debug {
									log.Print("Broadcasting to " + connector.ID + " (type:" + connector.Type + ") for route " + route.Connectors)
									log.Printf("Message: %+v", message)
									log.Print("")
								}
								for _, target := range strings.Split(route.Targets, ",") {
									if target == "*" {
										target = message.In.Target
									}
									c := connectors.MakeConnector(connector.Type).(connectors.Connector)
									c.Publish(connector, message, target)
								}
							}
						}
					}
				}
			}
		}
	}
}

func matchRoute(pmatch bool, mValue string, rValue string) (match bool) {
	match = false
	if pmatch {
		if mValue == rValue {
			match = true
		} else {
			if rValue == "*" {
				match = true
			}
		}
	}
	return match
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
