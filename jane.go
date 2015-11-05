package main

import (
	"github.com/mmcquillan/jane/listeners"
	"github.com/mmcquillan/jane/models"
	"github.com/yvasiyarov/gorelic"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	config := models.Load()
	models.Flags(&config)
	models.Logging(&config)
	runProfiling(&config)
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
			case "monitor":
				go listeners.Monitor(config, listener)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func runProfiling(config *models.Config) {
	if config.NewRelic != "" {
		agent := gorelic.NewAgent()
		agent.NewrelicLicense = config.NewRelic
		agent.NewrelicName = config.Name
		agent.Verbose = config.Debug
		agent.Run()
	}
}
