package main

import (
	"github.com/mmcquillan/jane/listeners"
	"github.com/mmcquillan/jane/models"
	"sync"
)

var wg sync.WaitGroup

func main() {
	config := models.Load()
	models.Flags(&config)
	models.Logging(&config)
	wg.Add(len(config.Listeners))
	go runListener(&config)
	wg.Wait()
}

func runListener(config *models.Config) {
	for _, listener := range config.Listeners {
		if listener.Active {
			defer wg.Done()
			switch listener.Type {
			case "slack":
				go listeners.Slack(config, listener)
			case "cli":
				go listeners.Cli(config, listener)
			case "rss":
				go listeners.Rss(config, listener)
			}
		}
	}
}
