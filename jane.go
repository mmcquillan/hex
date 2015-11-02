package main

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/listeners"
	"github.com/mmcquillan/jane/relays"
	"sync"
)

var wg sync.WaitGroup

func main() {
	config := configs.Load()
	configs.Flags(&config)
	configs.Logging(&config)
	wg.Add(3)
	go runRelays(&config)
	go runListener(&config)
	go listenLoop(&config)
	wg.Wait()
}

func runRelays(config *configs.Config) {
	defer wg.Done()
	for _, relay := range config.Relays {
		if relay.Active {
			switch relay.Type {
			case "cli":
				go relays.CliIn(config, relay)
			case "slack":
				go relays.SlackIn(config, relay)
			}
		}
	}
}

func runListener(config *configs.Config) {
	defer wg.Done()
	for _, listener := range config.Listeners {
		if listener.Active {
			switch listener.Type {
			case "rss":
				go listeners.Rss(config, listener)
			}
		}
	}
}

func listenLoop(config *configs.Config) {
	defer wg.Done()
	listeners.Listen(config)
}
