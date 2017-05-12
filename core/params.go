package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/hexbotio/hex/models"
)

func Params(config *models.Config) {

	// capture the flag
	configFile := flag.String("config", "", "Location of the config file")
	validate := flag.Bool("validate", false, "Validate the config file")
	showVersion := flag.Bool("version", false, "Version of HexBot")
	flag.Parse()

	// set the config values
	config.ConfigFile = *configFile
	config.Validate = *validate

	// operate on any flags
	if *showVersion {
		fmt.Print("HexBot " + config.Version + "\n")
		os.Exit(0)
	}

}
