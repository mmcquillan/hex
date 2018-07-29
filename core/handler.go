package core

import (
	"os"
	"os/signal"

	"github.com/mmcquillan/hex/models"
)

func Handler(plugins *map[string]models.Plugin, config models.Config) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	go func(signalChannel chan os.Signal) {
		sig := <-signalChannel
		config.Logger.Info("Received Signal " + sig.String())
		config.Logger.Info("Stopping Plugins...")
		StopPlugins(*plugins, config)
		config.Logger.Info("Stopping Hex Bot...")
		os.Exit(0)
	}(signalChannel)
}
