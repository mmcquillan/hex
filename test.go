package main

import (
	"bitbucket.org/prysm/devops-robot/configs"
	"bitbucket.org/prysm/devops-robot/listeners"
	"fmt"
	"os"
)

func main() {
	config := loadConfig()
	_, messages := listeners.Bamboo(&config, "")
	for _, m := range messages {
		fmt.Println(m.Description)
	}
}

func loadConfig() (config configs.Config) {
	configFile := "/home/mmcquillan/slacker.config"
	if configs.CheckConfig(configFile) {
		config = configs.ReadConfig(configFile)
	} else {
		configs.WriteDefaultConfig(configFile)
		fmt.Println("Please configure " + configFile)
		os.Exit(1)
	}
	return config
}
