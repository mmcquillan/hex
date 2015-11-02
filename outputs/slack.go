package outputs

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/nlopes/slack"
)

func Slack(config *configs.Config, message Message) {
	api := slack.New(config.SlackToken)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = config.Name
	params.IconEmoji = config.JaneEmoji
	if message.Description != "" {
		color := ColorMe(message.Status)
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

func ColorMe(status string) (color string) {
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
