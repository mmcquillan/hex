package services

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

// Input Inputs to slack messages in channels Jane is present in
func (x Slack) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	api := slack.New(connector.Key)
	api.SetDebug(connector.Debug)

	// get channel array
	channels := make(map[string]string)
	channelList, err := api.GetChannels(true)
	if err != nil {
		log.Printf("%s\n", err)
	}
	for _, channel := range channelList {
		channels[channel.ID] = "#" + channel.Name
	}

	// get user array
	users := make(map[string]string)
	userList, err := api.GetUsers()
	if err != nil {
		log.Printf("%s\n", err)
	}
	for _, user := range userList {
		users[user.ID] = user.Name
	}

	// listen to messages
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
					var found bool
					m.In.ConnectorType = connector.Type
					m.In.ConnectorID = connector.ID
					m.In.Tags = connector.Tags
					m.In.Target, found = channels[ev.Channel]
					if !found {
						m.In.Target = ev.Channel
					}
					m.In.User = users[ev.User]
					m.In.Text = html.UnescapeString(ev.Text)
					m.In.Process = true
					inputMsgs <- m

				}
			}
		}
	}
}

// Action Not implemented
func (x Slack) Action(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

// Output Outputes messages to slack
func (x Slack) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	api := slack.New(connector.Key)
	for {
		message := <-outputMsgs
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
			if target == "*" {
				target = message.In.Target
			}
			if target != "" && target != "*" {
				api.PostMessage(target, msg, params)
			}
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
