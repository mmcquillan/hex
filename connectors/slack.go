package connectors

import (
	"github.com/nlopes/slack"
	"github.com/projectjane/jane/models"
	"log"
)

type Slack struct {
	Connector models.Connector
}

func (x Slack) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
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

					var r []models.Route
					r = append(r, models.Route{Match: "*", Connectors: connector.ID, Target: ev.Channel})
					for _, cr := range connector.Routes {
						r = append(r, cr)
					}

					var m models.Message
					m.Routes = r
					m.In.Source = connector.ID
					m.In.User = ev.User
					m.In.Text = ev.Text
					m.In.Process = true
					commandMsgs <- m

				}
			}
		}
	}
}

func (x Slack) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Slack) Publish(connector models.Connector, message models.Message, target string) {
	api := slack.New(connector.Key)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = "jane"
	params.IconEmoji = connector.Image
	if target == "" {
		target = "#general"
	}
	if message.Out.Detail != "" {
		color := slackColorMe(message.Out.Status)
		attachment := slack.Attachment{
			Title:     message.Out.Text,
			TitleLink: message.Out.Link,
			Text:      message.Out.Detail,
			Color:     color,
		}
		params.Attachments = []slack.Attachment{attachment}
	} else {
		msg = message.Out.Text
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
