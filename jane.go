package main

import (
	"github.com/mmcquillan/jane/listeners"
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
	wg.Add(activeListeners(&config))
	runListener(&config)
	defer wg.Done()
	wg.Wait()
}

func activeListeners(config *models.Config) (cnt int) {
	cnt = 0
	for _, listener := range config.Listeners {
		if listener.Active {
			cnt += 1
		}
	}
	if config.Debug {
		log.Print("Active Listner count: " + string(cnt))
	}
	return cnt
}

func runListener(config *models.Config) {
	for _, listener := range config.Listeners {
		if listener.Active {
			log.Print("Initializing " + listener.Name + " (" + listener.Type + ")")
			switch listener.Type {
			case "slack":
				go listeners.Slack(config, listener)
			case "cli":
				go listeners.Cli(config, listener)
			case "rss":
				go listeners.Rss(config, listener)
			case "monitor":
				go listeners.Monitor(config, listener)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
