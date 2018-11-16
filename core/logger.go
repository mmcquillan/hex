package core

import (
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/mmcquillan/hex/models"
)

// Logger function
func Logger(config *models.Config) {
	logOptions := hclog.LoggerOptions{
		Name:  config.BotName,
		Level: hclog.Info,
	}
	logLevel := strings.ToUpper(config.LogLevel)
	if logLevel == "ERROR" {
		logOptions.Level = hclog.Error
	}
	if logLevel == "DEBUG" {
		logOptions.Level = hclog.Debug
	}
	if logLevel == "TRACE" {
		logOptions.Level = hclog.Trace
	}
	if config.LogFile != "" {
		if _, err := os.Stat(config.LogFile); os.IsNotExist(err) {
			nf, err := os.Create(config.LogFile)
			if err != nil {
				log.Fatal("ERROR: Cannot create log file at "+config.LogFile, err)
			}
			nf.Close()
		}
		f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			log.Fatal("ERROR: Cannot get to log file at "+config.LogFile, err)
		}
		logOptions.Output = f
	}
	config.Logger = hclog.New(&logOptions)
	config.Logger.Info("Starting Hex (" + config.Version + ")")
	config.Logger.Info("Initializing Logger")
}
