package listeners

import (
	"bitbucket.org/prysm/devops-robot/configs"
	"github.com/nlopes/slack"
	"strconv"
	"time"
)

type Message struct {
	Destination string
	Title       string
	Description string
	Link        string
	Status      string
}

func Listen(config *configs.Config) {

	// init
	now := time.Now()
	messages := make([]Message, 0)
	bambooMarker := now.UTC().String()
	deployMarker := strconv.FormatInt(now.Unix(), 10) + "000"

	// general loop
	for {

		// bamboo
		bambooMarker, messages = Bamboo(config, bambooMarker)
		for _, m := range messages {
			Talk(config, m)
		}

		// deploys
		deployMarker, messages = Deploys(config, deployMarker)
		for _, m := range messages {
			Talk(config, m)
		}

		// wait
		time.Sleep(120 * time.Second)

	}

}

func Talk(config *configs.Config, message Message) {
	api := slack.New(config.SlackToken)
	params := slack.NewPostMessageParameters()
	params.Username = config.SlackerName
	params.IconEmoji = config.SlackerEmoji
	color := ColorMe(message.Status)
	attachment := slack.Attachment{
		Title:     message.Title,
		TitleLink: message.Link,
		Text:      message.Description,
		Color:     color,
	}
	params.Attachments = []slack.Attachment{attachment}
	api.PostMessage(message.Destination, "", params)
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
	default:
		color = "#DDDDDD"
	}
	return color
}
