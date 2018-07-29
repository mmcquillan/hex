package core

import (
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/mmcquillan/hex/models"
)

func Logging(config *models.Config) {
	logOptions := hclog.LoggerOptions{
		Name:  "hex",
		Level: hclog.Info,
	}
	if config.Quiet {
		logOptions.Level = hclog.Error
	}
	if config.Debug {
		logOptions.Level = hclog.Debug
	}
	if config.Trace {
		logOptions.Level = hclog.Trace
	}
	if config.LogFile != "" {
		if !FileExists(config.LogFile) {
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
	config.Logger.Info(". . .")
	config.Logger.Info("Starting HexBot " + config.Version)
	config.Logger.Info("http://hexbot.io")
}
