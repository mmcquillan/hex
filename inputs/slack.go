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
		config.Logger.Error("Slack Channel List" + " - " + err.Error())
	}
	for _, channel := range channelList {
		channels[channel.ID] = "#" + channel.Name
	}

	// decode the bot name
	var botName = config.BotName
	var useAt = false
	if strings.HasPrefix(config.BotName, "@") {
		useAt = true
		botName = strings.Replace(botName, "@", "", 1)
	}

	// get user array
	users := make(map[string]string)
	userList, err := api.GetUsers()
	if err != nil {
		config.Logger.Error("Slack User List" + " - " + err.Error())
	}
	for _, user := range userList {
		users[user.ID] = user.Name
		if botName == user.Name {
			botName = user.ID
		}
	}

	// validate bot
	if botName == "" {
		config.Logger.Warn("Bot name is not the same as configured in Slack")
	} else {
		botName = "<@" + botName + ">"
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

					var match = false
					var input = ""
					if useAt && strings.HasPrefix(ev.Text, botName) {
						match = true
						input = strings.Replace(html.UnescapeString(ev.Text), botName+" ", "", 1)
					}
					if !useAt && strings.HasPrefix(ev.Text, config.BotName) {
						match = true
						input = strings.Replace(html.UnescapeString(ev.Text), config.BotName+" ", "", 1)
					}
					if match {
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
