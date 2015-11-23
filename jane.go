package main

import (
	"github.com/projectjane/jane/connectors"
	"github.com/projectjane/jane/models"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Received signal %v", sig)
			log.Print("Stopping jane bot...")
			os.Exit(0)
		}
	}()
	config := models.Load()
	models.Logging(&config)
	log.Print("---")
	log.Print("Starting jane bot...")
	wg.Add(activeConnectors(&config))
	runConnector(&config)
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

func runConnector(config *models.Config) {
	for _, connector := range config.Connectors {
		if connector.Active {
			log.Print("Initializing " + connector.ID + " (type: " + connector.Type + ")")
			c := connectors.MakeConnector(connector.Type).(connectors.Connector)
			go c.Listen(config, connector)
			time.Sleep(2 * time.Second)
		}
	}
}
