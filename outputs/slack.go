package outputs

import (
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/nlopes/slack"
)

// Slack struct
type Slack struct {
}

// Output function
func (x Slack) Write(outputMsgs <-chan models.Message, service models.Service) {
	api := slack.New(service.Config["Key"])
	for {
		message := <-outputMsgs
		msg := ""
		params := slack.NewPostMessageParameters()
		params.Username = service.BotName
		image := service.Config["Image"]
		if image == "" {
			image = ":nut_and_bolt:"
		}
		params.IconEmoji = image
		if !message.Success {
			attachment := slack.Attachment{
				Title:      strings.Join(message.Response[:], "\n"),
				Color:      "danger",
				MarkdownIn: []string{"text"},
			}
			params.Attachments = []slack.Attachment{attachment}
		} else if message.Inputs["hex.pipeline.alert"] == "true" && message.Success {
			attachment := slack.Attachment{
				Title:      strings.Join(message.Response[:], "\n"),
				Color:      "good",
				MarkdownIn: []string{"text"},
			}
			params.Attachments = []slack.Attachment{attachment}
		} else {
			msg = strings.Join(message.Response[:], "\n")
		}
		api.PostMessage(message.Inputs["hex.output"], msg, params)
	}
}
