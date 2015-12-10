package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"log"
	"strings"
)

func Publishers(publishMsgs <-chan models.Message, config *models.Config) {
	log.Print("Initializing Publishers")
	for {
		message := <-publishMsgs
		for _, route := range message.Routes {
			if strings.Contains(message.Out.Text, route.Match) || route.Match == "*" {
				for _, connector := range config.Connectors {
					if connector.Active {
						if sendToConnector(connector.ID, route.Connectors) {
							if connector.Debug {
								log.Print("Broadcasting to " + connector.ID + " (type:" + connector.Type + ") for route " + route.Connectors)
								log.Printf("Message: %+v", message)
								log.Print("")
							}
							c := connectors.MakeConnector(connector.Type).(connectors.Connector)
							c.Publish(connector, message, route.Target)
						}
					}
				}
			}
		}
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
