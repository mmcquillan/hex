package relays

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/nlopes/slack"
)

func SlackIn(config *configs.Config, relay configs.Relay) {
	api := slack.New(relay.Resource)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {
					Parse(config, relay, ev.Channel, ev.Text)
				}
			}
		}
	}
}

func SlackOut(config *configs.Config, relay configs.Relay, message Message) {
	api := slack.New(relay.Resource)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = config.Name
	params.IconEmoji = relay.Image
	if message.Description != "" {
		color := SlackColorMe(message.Status)
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
	api.PostMessage(message.Destination, msg, params)
}

func SlackColorMe(status string) (color string) {
	switch status {
	case "Successful":
		color = "good"
	case "SUCCESS":
		color = "good"
	case "Failed":
		color = "danger"
	case "FAILED":
		color = "danger"
	case "FAIL":
		color = "danger"
	case "NONE":
		color = "#DDDDDD"
	default:
		color = "#DDDDDD"
	}
	return color
}
