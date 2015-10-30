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
	wg.Add(2)
	go commandLoop(&config)
	go listenLoop(&config)
	wg.Wait()
}

func commandLoop(config *configs.Config) {
	defer wg.Done()
	inputs.Input(config)
}

func listenLoop(config *configs.Config) {
	defer wg.Done()
	listeners.Listen(config)
}
