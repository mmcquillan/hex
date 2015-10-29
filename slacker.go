package main

import (
	"bitbucket.org/prysm/devops-robot/commands"
	"bitbucket.org/prysm/devops-robot/configs"
	"bitbucket.org/prysm/devops-robot/listeners"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	configFile := flag.String("config", "/etc/slacker.config", "Config file")
	logFile := flag.String("log", "/var/log/slacker.log", "Log file")
	flag.Parse()
	setLogs(*logFile)
	config := loadConfig(*configFile)
	wg.Add(2)
	go commandLoop(&config)
	go listenLoop(&config)
	wg.Wait()
}

func setLogs(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Cannot find log file at " + logFile)
		panic(err)
	}
	log.SetOutput(f)
	log.Println("Starting the slacker bot")
}

func loadConfig(configFile string) (config configs.Config) {
	if configs.CheckConfig(configFile) {
		config = configs.ReadConfig(configFile)
	} else {
		configs.WriteDefaultConfig(configFile)
		fmt.Println("Missing config file, creating...")
		fmt.Println("Please configure " + configFile)
		os.Exit(1)
	}
	return config
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
