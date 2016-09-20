package core

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
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
			match := true
			if _, chk := chans[route.Publish.ConnectorID]; !chk {
				match = false
			}
			matchRouteFull(&match, message.Out.Text+" "+message.Out.Detail, route.Match.Message)
			matchRouteSimple(&match, message.In.ConnectorType, route.Match.ConnectorType)
			matchRouteSimple(&match, message.In.ConnectorID, route.Match.ConnectorID)
			matchRouteTags(&match, message.In.Tags, route.Match.Tags)
			matchRouteSimple(&match, message.In.Target, route.Match.Target)
			matchRouteSimple(&match, message.In.User, route.Match.User)
			if match {
				message.Out.Target = route.Publish.Target
				chans[route.Publish.ConnectorID] <- message
			}
		}
	}
}

func matchRouteFull(match *bool, mValue string, rValue string) {
	if *match {
		*match, _ = parse.Match(rValue, mValue)
	}
}

func matchRouteSimple(match *bool, mValue string, rValue string) {
	if *match {
		*match = parse.SimpleMatch(mValue, rValue)
	}
}

func matchRouteTags(match *bool, mValue string, rValue string) {
	if *match {
		*match = parse.TagMatch(mValue, rValue)
	}
}
