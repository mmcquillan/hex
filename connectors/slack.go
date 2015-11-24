package connectors

import (
	"github.com/nlopes/slack"
	"github.com/projectjane/jane/commands"
	"github.com/projectjane/jane/models"
	"log"
	"regexp"
	"strings"
)

type Slack struct {
	Connector models.Connector
}

func (x Slack) Listen(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	api := slack.New(connector.Key)
	api.SetDebug(connector.Debug)
	rtm := api.NewRTM()
	if connector.Debug {
		log.Print("Starting slack websocket api for " + connector.ID)
	}
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

					if connector.Debug {
						log.Print("Evaluating incoming slack message")
					}

					var process = true

					// make sure they are talking to and not about us
					var msg = ev.Text
					tokmsg := strings.Split(strings.TrimSpace(msg), " ")
					if strings.ToLower(tokmsg[0]) != "jane" {
						process = false

						var jiraRegex = regexp.MustCompile(`(?i)(SYN|MED|STD)-[0-9]+`)
						matches := jiraRegex.FindAllString(msg, -1)

						if len(matches) > 0 {

							var r []models.Route
							r = append(r, models.Route{Match: "*", Connectors: connector.ID, Target: ev.Channel})
							for _, cr := range connector.Routes {
								r = append(r, cr)
							}

							for _, match := range matches {
								m := models.Message{
									Routes:      r,
									Source:      ev.User,
									Request:     "jira " + match,
									Title:       "",
									Description: "",
									Link:        "",
									Status:      "",
								}

								commands.Parse(config, &m)
								Broadcast(config, m)
							}

							return
						}
					}

					// remove me from the request and clean
					msg = strings.Replace(msg, tokmsg[0], "", 1)
					msg = strings.TrimSpace(msg)

					// see if nothing is said
					if msg == "" {
						process = false
					}

					if process {
						if connector.Debug {
							log.Print("Processing incoming slack message")
						}
						var r []models.Route
						r = append(r, models.Route{Match: "*", Connectors: connector.ID, Target: ev.Channel})
						for _, cr := range connector.Routes {
							r = append(r, cr)
						}
						m := models.Message{
							Routes:      r,
							Source:      ev.User,
							Request:     msg,
							Title:       "",
							Description: "",
							Link:        "",
							Status:      "",
						}
						commands.Parse(config, &m)
						Broadcast(config, m)
					}
				}
			}
		}
	}
}

func (x Slack) Command(config *models.Config, message *models.Message) {
	return
}

func (x Slack) Publish(config *models.Config, connector models.Connector, message models.Message, target string) {
	api := slack.New(connector.Key)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = "jane"
	params.IconEmoji = connector.Image
	if target == "" {
		target = "#general"
	}
	if message.Description != "" {
		color := slackColorMe(message.Status)
		attachment := slack.Attachment{
			Title:     message.Title,
			TitleLink: message.Link,
			Text:      message.Description,
			Color:     color,
		}
		params.Attachments = []slack.Attachment{attachment}
	} else {
		msg = message.Title
	}
	api.PostMessage(target, msg, params)
}

func slackColorMe(status string) (color string) {
	switch status {
	case "SUCCESS":
		color = "good"
	case "WARN":
		color = "warning"
	case "FAIL":
		color = "danger"
	case "NONE":
		color = "#DDDDDD"
	default:
		color = "#DDDDDD"
	}
	return color
}
