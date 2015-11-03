package relays

import (
	"github.com/mmcquillan/jane/models"
	"github.com/nlopes/slack"
)

func Slack(config *models.Config, relay models.Relay, message models.Message) {
	api := slack.New(relay.Resource)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = config.Name
	params.IconEmoji = relay.Image
	if message.Target == "" {
		message.Target = "#general"
	}
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
	api.PostMessage(message.Target, msg, params)
}

func SlackColorMe(status string) (color string) {
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
