package main

import (
	"github.com/projectjane/jane/core"
	"github.com/projectjane/jane/models"
	"sync"
)

func main() {
	core.CatchExit()
	config := core.LoadConfig()
	core.Logging(&config)
	commandMsgs := make(chan models.Message, 1)
	publishMsgs := make(chan models.Message, 1)
	var wg sync.WaitGroup
	wg.Add(core.ActiveConnectors(&config))
	wg.Add(3)
	go core.Listeners(commandMsgs, &config)
	go core.Commands(commandMsgs, publishMsgs, &config)
	go core.Publishers(publishMsgs, &config)
	defer wg.Done()
	wg.Wait()
}
