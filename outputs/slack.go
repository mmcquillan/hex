package outputs

import (
	"strings"

	"github.com/nlopes/slack"
	"github.com/hexbotio/hex/models"
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
		params.IconEmoji = service.Config["Image"]
		if !message.Success {
			attachment := slack.Attachment{
				Title:      strings.Join(message.Response[:], "\n"),
				Color:      "danger",
				MarkdownIn: []string{"text"},
			}
			params.Attachments = []slack.Attachment{attachment}
		} else {
			msg = strings.Join(message.Response[:], "\n")
		}
		api.PostMessage(message.Inputs["hex.output"], msg, params)
	}
}
