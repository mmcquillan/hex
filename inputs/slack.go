package inputs

import (
	"html"
	"log"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/nlopes/slack"
)

// Slack struct
type Slack struct {
}

// Read function
func (x Slack) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
	var useAt = false
	if strings.ToLower(service.Config["UseAt"]) == "true" {
		useAt = true
	}
	api := slack.New(service.Config["Key"])
	var slackDebug = false
	if strings.ToLower(service.Config["SlackDebug"]) == "true" {
		slackDebug = true
	}
	api.SetDebug(slackDebug)

	// get channel array
	channels := make(map[string]string)
	channelList, err := api.GetChannels(true)
	if err != nil {
		log.Printf("%s\n", err)
	}
	for _, channel := range channelList {
		channels[channel.ID] = "#" + channel.Name
	}

	// get user array
	bot := ""
	users := make(map[string]string)
	userList, err := api.GetUsers()
	if err != nil {
		log.Printf("%s\n", err)
	}
	for _, user := range userList {
		users[user.ID] = user.Name
		if service.BotName == user.Name {
			bot = user.ID
		}
	}

	// validate bot
	if bot == "" {
		log.Print("WARNING - Bot does not seem to be configured in Slack as '" + service.BotName + "'")
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

					if useAt && strings.HasPrefix(ev.Text, bot) {
						input := strings.Replace(html.UnescapeString(ev.Text), bot+" ", "", 1)
						message := models.MakeMessage(service.Type, service.Name, channel, users[ev.User], input)
						inputMsgs <- message
					}
					if !useAt && strings.HasPrefix(ev.Text, service.BotName) {
						input := strings.Replace(html.UnescapeString(ev.Text), service.BotName+" ", "", 1)
						message := models.MakeMessage(service.Type, service.Name, channel, users[ev.User], input)
						inputMsgs <- message
					}
				}
			}
		}
	}
}
