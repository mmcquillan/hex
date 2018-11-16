package main

import (
	"sync"

	"github.com/mmcquillan/hex/core"
	"github.com/mmcquillan/hex/models"
)

var version string

func main() {

	// initialize global vars
	var wg sync.WaitGroup
	var config models.Config
	var rules = make(map[string]models.Rule)
	var plugins = make(map[string]models.Plugin)
	var inputMsgs = make(chan models.Message, 1)
	var outputMsgs = make(chan models.Message, 1)

	// initialize app
	config = core.Config(version)
	core.Logger(&config)
	core.Handler(&plugins, config)
	core.Rules(&rules, config)
	core.Plugins(&plugins, rules, config)

	// run application
	wg.Add(3)
	go core.Inputs(inputMsgs, &rules, config)
	go core.Matcher(inputMsgs, outputMsgs, &plugins, &rules, config)
	go core.Outputs(outputMsgs, &plugins, config)

	// run indefinately
	defer wg.Done()
	wg.Wait()

}
