package main

import (
	"github.com/projectjane/jane/api"
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
	api.Start()
	defer wg.Done()
	wg.Wait()
}
