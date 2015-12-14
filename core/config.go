package core

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"github.com/projectjane/jane/models"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig(params models.Params) (config models.Config) {
	configFile := locateConfig(params)
	if checkConfig(configFile) {
		config = readConfig(configFile)
		if params.Validate {
			fmt.Println("SUCCESS - Config file is valid: " + configFile)
			os.Exit(0)
		}
	} else {
		os.Exit(1)
	}
	return config
}

func locateConfig(params models.Params) (configFile string) {
	file := "jane.json"

	zero := params.ConfigFile
	if FileExists(zero) {
		return zero
	}

	first, _ := osext.ExecutableFolder()
	first += "/" + file
	if FileExists(first) {
		return first
	}

	second, _ := homedir.Dir()
	second += "/" + file
	if FileExists(second) {
		return second
	}

	third := "/etc/" + file
	if FileExists(third) {
		return third
	}

	return file

}

func readConfig(location string) (config models.Config) {
	file, err := ioutil.ReadFile(location)
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Print(err)
	}
	return config
}

func checkConfig(location string) (exists bool) {
	exists = true
	if _, err := os.Stat(location); os.IsNotExist(err) {
		fmt.Println("Error - Missing a config file")
		exists = false
	}
	if exists {
		file, err := ioutil.ReadFile(location)
		if err != nil {
			log.Print(err)
		}
		var js interface{}
		exists = json.Unmarshal(file, &js) == nil
		if !exists {
			fmt.Println("Error - Config file does not appear to be valid json")
		}
	}
	return exists
}
