package core

import (
	"os"
	"os/signal"

	"github.com/hexbotio/hex/models"
)

func Handler(plugins *map[string]models.Plugin, config models.Config) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			config.Logger.Info("Received Interrupt Signal")
			config.Logger.Info("Stopping Plugins...")
			StopPlugins(*plugins, config)
			config.Logger.Info("Stopping Hex Bot...")
			os.Exit(0)
		}
	}()
}
