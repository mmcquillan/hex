package inputs

import (
	"fmt"
	"html"
	"regexp"
	"strings"
	"time"

	"github.com/mmcquillan/hex/models"
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

	// initialize
	x.API = slack.New(config.SlackToken)
	x.API.SetDebug(config.SlackDebug)
	x.Users = make(map[string]string)
	x.Channels = make(map[string]string)

	// get channel array
	x.Channels = x.updateChannels(config)
	go x.refreshChannels(config)

	// listen to messages
	rtm := x.API.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {

					// resolve channel
					channel, found := x.Channels[ev.Channel]
					if !found {
						channel = ev.Channel
					}

					// unencode slack input
					var input = html.UnescapeString(ev.Text)
					config.Logger.Trace("Slack Input - Raw Input: " + input)

					// resolve slack users
					re := regexp.MustCompile("<@([A-Za-z0-9]+)>")
					tokens := re.FindAllString(input, -1)
					if len(tokens) > 0 {
						for _, t := range tokens {
							input = strings.Replace(input, t, "@"+x.resolveUser(t, config), -1)
						}
					}

					message := models.NewMessage()
					message.Attributes["hex.botname"] = config.BotName
					message.Attributes["hex.service"] = "slack"
					message.Attributes["hex.channel"] = channel
					message.Attributes["hex.channel.id"] = ev.Channel
					message.Attributes["hex.user"] = x.resolveUser(ev.User, config)
					message.Attributes["hex.input"] = input
					config.Logger.Debug("Slack Input - ID:" + message.Attributes["hex.id"])
					config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
					inputMsgs <- message

				}
			}
		}
	}
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

func (x Slack) resolveUser(user string, config models.Config) (resolved string) {
	if strings.HasPrefix(user, "<@") && strings.HasSuffix(user, ">") {
		user = string(user[2 : len(user)-1])
	}
	if u, chk := x.Users[user]; chk {
		return u
	}
	u, err := x.API.GetUserInfo(user)
	if err != nil {
		config.Logger.Error("Slack Input - GetUserInfo Error: " + err.Error())
	} else {
		resolved = u.Name
		x.Users[user] = resolved
	}
	return resolved
}
