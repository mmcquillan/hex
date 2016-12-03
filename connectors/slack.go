package connectors

import (
	"html"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"github.com/projectjane/jane/models"
)

// Slack Represents Slack connector
type Slack struct {
	Connector models.Connector
}

// Listen Listens to slack messages in channels Jane is present in
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
					m.In.Tags = connector.Tags
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

// Command Not implemented
func (x Slack) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

// Publish Publishes messages to slack
func (x Slack) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	api := slack.New(connector.KeyValues["Key"])
	for {
		message := <-publishMsgs
		msg := ""
		params := slack.NewPostMessageParameters()
		params.Username = connector.BotName
		params.IconEmoji = connector.Image
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
		for _, target := range strings.Split(message.Out.Target, ",") {
			if target == "" {
				target = "#general"
			}
			if target == "*" {
				target = message.In.Target
			}
			api.PostMessage(target, msg, params)
		}
	}
}

// Help Not Implemented
func (x Slack) Help(connector models.Connector) (help []string) {
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
