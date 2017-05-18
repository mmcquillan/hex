package main

import (
	"sync"

	"github.com/hexbotio/hex/core"
	"github.com/hexbotio/hex/models"
)

var version string

func main() {
	var wg sync.WaitGroup
	var config models.Config
	config.Version = version
	core.Params(&config)
	core.Config(&config)
	core.Starter(&config)
	inputMsgs := make(chan models.Message, 1)
	outputMsgs := make(chan models.Message, 1)
	wg.Add(3)
	go core.Inputs(inputMsgs, &config)
	go core.Pipeline(inputMsgs, outputMsgs, &config)
	go core.Outputs(outputMsgs, &config)
	defer wg.Done()
	wg.Wait()
}
