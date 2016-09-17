package connectors

import (
	"html"
	"log"

	"github.com/nlopes/slack"
	"github.com/projectjane/jane/models"
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

					var m models.Message
					m.In.ConnectorType = connector.Type
					m.In.ConnectorID = connector.ID
					m.In.Target = ev.Channel
					m.In.User = ev.User
					m.In.Text = html.UnescapeString(ev.Text)
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
			Title:      message.Out.Text,
			TitleLink:  message.Out.Link,
			Text:       message.Out.Detail,
			Color:      color,
			MarkdownIn: []string{"text"},
		}
		params.Attachments = []slack.Attachment{attachment}
	} else {
		msg = message.Out.Text
	}
	api.PostMessage(target, msg, params)
}

func (x Slack) Help(connector models.Connector) (help string) {
	return
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
