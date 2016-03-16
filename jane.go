package main

import (
	"github.com/projectjane/jane/core"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/api"
	"sync"
)

func main() {
	core.CatchExit()
	params := core.LoadParams()
	config := core.LoadConfig(params)
	core.Logging(&config)
	commandMsgs := make(chan models.Message, 1)
	publishMsgs := make(chan models.Message, 1)
	api.StartRestServer(commandMsgs, publishMsgs)
	var wg sync.WaitGroup
	wg.Add(core.ActiveConnectors(&config))
	wg.Add(3)
	go core.Listeners(commandMsgs, &config)
	go core.Commands(commandMsgs, publishMsgs, &config)
	go core.Publishers(publishMsgs, &config)
	defer wg.Done()
	wg.Wait()
}
