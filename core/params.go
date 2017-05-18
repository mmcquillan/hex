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
	startup := flag.Bool("startup", false, "Startup file for HexBot")
	flag.Parse()

	// set the config values
	config.ConfigFile = *configFile
	config.Validate = *validate

	// operate on any flags
	if *showVersion {
		fmt.Print("HexBot " + config.Version + "\n")
		os.Exit(0)
	}

	// startup file
	if *startup {
		fmt.Print("\n")
		fmt.Print("1. Create file: /etc/systemd/system/hex.service\n")
		fmt.Print("--------------------------------\n")
		fmt.Print("[Unit]\n")
		fmt.Print("Description=HexBot\n")
		fmt.Print("\n")
		fmt.Print("[Service]\n")
		fmt.Print("Type=simple\n")
		fmt.Print("ExecStart=/usr/local/bin/hex\n")
		fmt.Print("Restart=on-failure\n")
		fmt.Print("--------------------------------\n")
		fmt.Print("\n")
		fmt.Print("2. Run: systemctl enable hex.service\n")
		fmt.Print("\n")
		os.Exit(0)
	}

}
