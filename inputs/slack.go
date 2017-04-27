package inputs

import (
	"html"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"github.com/projectjane/jane/models"
)

// Slack struct
type Slack struct {
}

// Read function
func (x Slack) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
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
	users := make(map[string]string)
	userList, err := api.GetUsers()
	if err != nil {
		log.Printf("%s\n", err)
	}
	for _, user := range userList {
		users[user.ID] = user.Name
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

					message := models.MakeMessage(service.Type, service.Name, channel, users[ev.User], html.UnescapeString(ev.Text))
					inputMsgs <- message

				}
			}
		}
	}
}
