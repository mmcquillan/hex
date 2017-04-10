package main

import (
	"github.com/projectjane/jane/core"
	"github.com/projectjane/jane/models"
	"sync"
)

var version string

func main() {
	core.CatchExit()
	params := core.Params()
	config := core.Config(params, version)
	core.Logging(&config)
	inputMsgs := make(chan models.Message, 1)
	outputMsgs := make(chan models.Message, 1)
	var wg sync.WaitGroup
	wg.Add(core.ActiveConnectors(&config))
	wg.Add(3)
	go core.Inputs(inputMsgs, &config)
	go core.Commands(inputMsgs, outputMsgs, &config)
	go core.Outputs(outputMsgs, &config)
	go core.Listeners(commandMsgs, &config)
	go core.Commands(commandMsgs, publishMsgs, &config)
	go core.Publishers(publishMsgs, &config)
	defer wg.Done()
	wg.Wait()
}
