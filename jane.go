package main

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/listeners"
	"github.com/nlopes/slack"
	"sync"
)

var wg sync.WaitGroup

func main() {
	config := configs.Load()
	configs.Logging(&config)
	wg.Add(2)
	go commandLoop(&config)
	go listenLoop(&config)
	wg.Wait()
}

func commandLoop(config *configs.Config) {
	defer wg.Done()
	api := slack.New(config.SlackToken)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if ev.User != "" {
					commands.Parse(config, ev.Channel, ev.Text)
				}
			}
		}
	}
}

func listenLoop(config *configs.Config) {
	defer wg.Done()
	listeners.Listen(config)
}
