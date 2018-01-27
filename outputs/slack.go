package outputs

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
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
	if rule.Threaded && message.Attributes["hex.slack.response"] != "" {
		params.ThreadTimestamp = message.Attributes["hex.slack.response"]
	}
	params.IconEmoji = image
	for _, output := range message.Outputs {
		if message.Debug && parse.EitherMember(config.Admins, message.Attributes["hex.user"], message.Attributes["hex.channel"]) {
			output.Response = output.Response + "\n\n[ Debug: " + output.Command + " ]"
		}
		if message.Attributes["hex.rule.format"] == "true" {
			msg = "*" + message.Attributes["hex.rule.name"] + "*"
			color := "grey"
			if output.Success {
				color = "good"
			} else {
				color = "danger"
			}
			attachment := slack.Attachment{
				Text:       "```" + output.Response + "```",
				Color:      color,
				MarkdownIn: []string{"text"},
			}
			params.Attachments = append(params.Attachments, attachment)
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
		out := fmt.Sprintf("\nMESSAGE DEBUG (%d sec to complete)\n", message.EndTime-message.StartTime)
		for _, key := range keys {
			if strings.HasPrefix(key, "hex.var.") {
				out = out + fmt.Sprintf("  %s: '%s'\n", key, "********")
			} else {
				out = out + fmt.Sprintf("  %s: '%s'\n", key, message.Attributes[key])
			}
		}
		attachment := slack.Attachment{
			Text:       "```" + out + "```",
			Color:      "grey",
			MarkdownIn: []string{"text"},
		}
		params.Attachments = append(params.Attachments, attachment)
	}
	if message.Attributes["hex.channel"] != "" {
		_, respTimestamp, err := api.PostMessage(message.Attributes["hex.channel"], msg, params)
		message.Attributes["hex.slack.response"] = respTimestamp
		if err != nil {
			config.Logger.Error("Slack Message Send Error: " + err.Error())
		}
	}
	if message.Attributes["hex.rule.channel"] != "" && message.Attributes["hex.channel"] != message.Attributes["hex.rule.channel"] {
		_, respTimestamp, err := api.PostMessage(message.Attributes["hex.rule.channel"], msg, params)
		message.Attributes["hex.slack.response"] = respTimestamp
		if err != nil {
			config.Logger.Error("Slack Message Send Error: " + err.Error())
		}
	}
}
