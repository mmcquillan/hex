package main

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/core"
	"github.com/projectjane/jane/models"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// catch sigterm
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Received signal %v", sig)
			log.Print("Stopping jane bot...")
			os.Exit(0)
		}
	}()

	// setup jane config
	config := models.Load()
	models.Logging(&config)
	log.Print("---")
	log.Print("Starting jane bot...")

	// setup channels and run
	commandMsgs := make(chan models.Message, 1)
	publishMsgs := make(chan models.Message, 1)
	wg.Add(activeConnectors(&config))
	wg.Add(3)
	go runListeners(commandMsgs, &config)
	go runCommands(commandMsgs, publishMsgs, &config)
	go runPublishers(publishMsgs, &config)
	defer wg.Done()
	wg.Wait()

}

func activeConnectors(config *models.Config) (cnt int) {
	cnt = 0
	for _, connector := range config.Connectors {
		if connector.Active {
			cnt += 1
		}
	}
	return cnt
}

func runListeners(commandMsgs chan<- models.Message, config *models.Config) {
	for _, connector := range config.Connectors {
		if connector.Active {
			log.Print("Initializing " + connector.ID + " listener (type: " + connector.Type + ")")
			c := connectors.MakeConnector(connector.Type).(connectors.Connector)
			go c.Listen(commandMsgs, connector)
		}
	}
}

func runCommands(commandMsgs <-chan models.Message, publishMsgs chan<- models.Message, config *models.Config) {
	log.Print("Initializing Commands")
	for {
		m := <-commandMsgs
		if m.In.Process {
			core.Commands(m, publishMsgs, config)
			for _, connector := range config.Connectors {
				if connector.Active {
					c := connectors.MakeConnector(connector.Type).(connectors.Connector)
					go c.Command(m, publishMsgs, connector)
				}
			}
		} else {
			publishMsgs <- m
		}
	}
}

func runPublishers(publishMsgs <-chan models.Message, config *models.Config) {
	log.Print("Initializing Publishers")
	for {
		message := <-publishMsgs
		for _, route := range message.Routes {
			if strings.Contains(message.Out.Text, route.Match) || route.Match == "*" {
				for _, connector := range config.Connectors {
					if connector.Active {
						if sendToConnector(connector.ID, route.Connectors) {
							if connector.Debug {
								log.Print("Broadcasting to " + connector.ID + " (type:" + connector.Type + ") for route " + route.Connectors)
								log.Printf("Message: %+v", message)
								log.Print("")
							}
							c := connectors.MakeConnector(connector.Type).(connectors.Connector)
							c.Publish(connector, message, route.Target)
						}
					}
				}
			}
		}
	}
}

func sendToConnector(connId string, connectors string) (send bool) {
	send = false
	if connectors == "*" {
		send = true
	}
	r := strings.Split(connectors, ",")
	for _, v := range r {
		if v == connId {
			send = true
		}
	}
	return send
}
