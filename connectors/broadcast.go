package connectors

import (
	"github.com/mmcquillan/jane/models"
	"log"
	"strings"
)

func Broadcast(config *models.Config, message models.Message) {
	for _, route := range message.Routes {
		if strings.Contains(message.Title, route.Match) || route.Match == "*" {
			for _, connector := range config.Connectors {
				if connector.Active {
					if sendToConnector(connector.ID, route.Connectors) {
						if config.Debug {
							log.Print("Broadcasting to " + connector.ID + " (" + connector.Type + ")")
						}
						c := MakeConnector(connector.Type).(Connector)
						c.Send(config, message, route.Target)
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
