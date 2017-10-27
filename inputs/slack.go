package inputs

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/nlopes/slack"
)

// Slack struct
type Slack struct {
	API      *slack.Client
	Channels map[string]string
	Users    map[string]string
}

// Read function
func (x Slack) Read(inputMsgs chan<- models.Message, config models.Config) {
	x.API = slack.New(config.SlackToken)
	x.API.SetDebug(config.SlackDebug)

	// get channel array
	x.Channels = x.updateChannels(config)
	go x.refreshChannels(config)

	// get user array
	x.Users = x.updateUsers(config)
	go x.refreshUsers(config)

	// decode the bot name
	var botName = config.BotName
	var useAt = false
	if strings.HasPrefix(config.BotName, "@") {
		useAt = true
		botName = strings.Replace(botName, "@", "", 1)
	}

	// convert the bot from name to id
	for userId, userName := range x.Users {
		if botName == userName {
			botName = userId
		}
	}

	// validate bot
	if botName == "" {
		config.Logger.Warn("Slack Input - Bot name is not the same as configured in Slack")
	} else {
		botName = "<@" + botName + ">"
	}

	// listen to messages
	rtm := x.API.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

					channel, found := x.Channels[ev.Channel]
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
						message.Attributes["hex.user"] = x.Users[ev.User]
						message.Attributes["hex.input"] = input
						message.Debug = debug
						config.Logger.Debug("Slack Input - ID:" + message.Attributes["hex.id"])
						config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
						inputMsgs <- message
					}

				}
			}
		}
	}
}

func (x Slack) updateUsers(config models.Config) map[string]string {
	users := make(map[string]string)
	userList, err := x.API.GetUsers()
	if err != nil {
		config.Logger.Error("Slack Input - User List " + err.Error())
	}
	for _, user := range userList {
		users[user.ID] = user.Name
	}
	return users
}

func (x Slack) refreshUsers(config models.Config) {
	time.Sleep(5 * time.Minute)
	x.Users = x.updateUsers(config)
}

func (x Slack) updateChannels(config models.Config) map[string]string {
	channels := make(map[string]string)
	channelList, err := x.API.GetChannels(true)
	if err != nil {
		config.Logger.Error("Slack Input - Channel List " + err.Error())
	}
	for _, channel := range channelList {
		channels[channel.ID] = "#" + channel.Name
	}
	return channels
}

func (x Slack) refreshChannels(config models.Config) {
	time.Sleep(5 * time.Minute)
	x.Channels = x.updateChannels(config)
}
