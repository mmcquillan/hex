package connectors

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"github.com/nlopes/slack"
	"log"
	"strings"
)

type Slack struct {
	Connector models.Connector
}

func (x Slack) Run(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	api := slack.New(connector.Key)
	api.SetDebug(config.Debug)
	rtm := api.NewRTM()
	if config.Debug {
		log.Print("Starting slack websocket api for " + connector.Name)
	}
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

					if config.Debug {
						log.Print("Evaluating incoming slack message")
					}
					var process = true

					// make sure they are talking to and not about us
					var msg = ev.Text
					tokmsg := strings.Split(strings.TrimSpace(msg), " ")
					if strings.ToLower(tokmsg[0]) != strings.ToLower(config.Name) {
						process = false
					}

					// remove me from the request and clean
					msg = strings.Replace(msg, tokmsg[0], "", 1)
					msg = strings.TrimSpace(msg)

					// see if nothing is said
					if msg == "" {
						process = false
					}

					if process {
						if config.Debug {
							log.Print("Processing incoming slack message")
						}
						for _, r := range connector.Routes {
							if strings.Contains(msg, r.Match) || r.Match == "*" {
								m := models.Message{
									Relays:      r.Relays,
									Target:      ev.Channel,
									Request:     msg,
									Title:       "",
									Description: "",
									Link:        "",
									Status:      "",
								}
								commands.Parse(config, m)
							}
						}
					}
				}
			}
		}
	}
}
