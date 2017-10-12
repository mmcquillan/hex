package outputs

import (
	"fmt"
	"sort"

	"github.com/hexbotio/hex/models"
	"github.com/nlopes/slack"
)

// Slack struct
type Slack struct {
}

// Output function
func (x Slack) Write(message models.Message, config models.Config) {
	api := slack.New(config.SlackToken)
	msg := ""
	params := slack.NewPostMessageParameters()
	params.Username = config.BotName
	image := config.SlackIcon
	if image == "" {
		image = ":nut_and_bolt:"
	}
	params.IconEmoji = image
	for _, output := range message.Outputs {
		if message.Attributes["hex.rule.format"] == "true" {
			color := "grey"
			if output.Success {
				color = "good"
			} else {
				color = "danger"
			}
			attachment := slack.Attachment{
				Title:      "TBD", //message.Attributes["hex.pipeline.name"],
				Text:       "```" + output.Response + "```",
				Color:      color,
				MarkdownIn: []string{"text"},
			}
			params.Attachments = []slack.Attachment{attachment}
		} else {
			msg = msg + output.Response + "\n"
		}
	}
	if message.Debug {
		keys := make([]string, 0, len(message.Attributes))
		for key, _ := range message.Attributes {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		msg = msg + fmt.Sprintf("\n```MESSAGE DEBUG (%d sec to complete)\n", models.MessageTimestamp()-message.CreateTime)
		for _, key := range keys {
			msg = msg + fmt.Sprintf("  %s: '%s'\n", key, message.Attributes[key])
		}
		msg = msg + "```"
	}
	api.PostMessage(message.Attributes["hex.channel"], msg, params)
}
