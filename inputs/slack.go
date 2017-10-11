package inputs

import (
	"html"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/nlopes/slack"
)

// Slack struct
type Slack struct {
}

// Read function
func (x Slack) Read(inputMsgs chan<- models.Message, config models.Config) {
	api := slack.New(config.SlackToken)
	api.SetDebug(config.SlackDebug)

	// get channel array
	channels := make(map[string]string)
	channelList, err := api.GetChannels(true)
	if err != nil {
		config.Logger.Error("Slack Channel List", err)
	}
	for _, channel := range channelList {
		channels[channel.ID] = "#" + channel.Name
	}

	// get user array
	bot := ""
	users := make(map[string]string)
	userList, err := api.GetUsers()
	if err != nil {
		config.Logger.Error("Slack User List", err)
	}
	for _, user := range userList {
		users[user.ID] = user.Name
		if config.BotName == user.Name {
			bot = user.ID
		}
	}

	// validate bot
	if bot == "" {
		config.Logger.Warn("Bot name is not the same as configured in Slack")
	} else {
		bot = "<@" + bot + ">"
	}

	// listen to messages
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

					channel, found := channels[ev.Channel]
					if !found {
						channel = ev.Channel
					}

					if strings.HasPrefix(ev.Text, bot) {
						input := strings.Replace(html.UnescapeString(ev.Text), bot+" ", "", 1)
						input, debug := parse.Flag(input, "--debug")
						message := models.NewMessage()
						message.Attributes["hex.service"] = "slack"
						message.Attributes["hex.channel"] = channel
						message.Attributes["hex.user"] = users[ev.User]
						message.Attributes["hex.input"] = input
						message.Debug = debug
						inputMsgs <- message
					}

				}
			}
		}
	}
}
