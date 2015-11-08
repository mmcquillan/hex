package main

import (
	"github.com/mmcquillan/jane/connectors"
	"github.com/mmcquillan/jane/models"
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
	wg.Add(activeconnectors(&config))
	runconnector(&config)
	defer wg.Done()
	wg.Wait()
}

func activeconnectors(config *models.Config) (cnt int) {
	cnt = 0
	for _, connector := range config.connectors {
		if connector.Active {
			cnt += 1
		}
	}
	if config.Debug {
		log.Print("Active Listner count: " + string(cnt))
	}
	return cnt
}

func runconnector(config *models.Config) {
	for _, connector := range config.connectors {
		if connector.Active {
			log.Print("Initializing " + connector.Name + " (" + connector.Type + ")")
			switch connector.Type {
			case "slack":
				go connectors.Slack(config, connector)
			case "cli":
				go connectors.Cli(config, connector)
			case "rss":
				go connectors.Rss(config, connector)
			case "monitor":
				go connectors.Monitor(config, connector)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
