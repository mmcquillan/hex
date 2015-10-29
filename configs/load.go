package configs

import (
	"fmt"
	"os"
)

func Load() (config Config) {

	configFile := Locate()
	if CheckConfig(configFile) {
		config = ReadConfig(configFile)
	} else {
		fmt.Println("Error - Missing a config file")
		os.Exit(1)
	}
	return config

}
