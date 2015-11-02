package main

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/inputs"
	"github.com/mmcquillan/jane/listeners"
	"sync"
)

var wg sync.WaitGroup

func main() {
	config := configs.Load()
	configs.Flags(&config)
	configs.Logging(&config)
	wg.Add(3)
	go runListener(&config)
	go commandLoop(&config)
	go listenLoop(&config)
	wg.Wait()
}

func runListener(config *configs.Config) {
	defer wg.Done()
	for _, listener := range config.Listeners {
		if listener.Active {
			switch listener.Type {
			case "rss":
				listeners.Rss(config, listener)
			}
		}
	}
}

func commandLoop(config *configs.Config) {
	defer wg.Done()
	inputs.Input(config)
}

func listenLoop(config *configs.Config) {
	defer wg.Done()
	listeners.Listen(config)
}
