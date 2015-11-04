package listeners

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"github.com/nlopes/slack"
	"strings"
)

func Slack(config *models.Config, listener models.Listener) {
	defer Recovery(config, listener)
	api := slack.New(listener.Resource)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

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
						for _, d := range listener.Destinations {
							if strings.Contains(msg, d.Match) || d.Match == "*" {
								m := models.Message{
									Relays:      d.Relays,
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
