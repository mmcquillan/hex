package core

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"github.com/projectjane/jane/models"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig() (config models.Config) {

	configFile := locateConfig()
	if checkConfig(configFile) {
		config = readConfig(configFile)
	} else {
		os.Exit(1)
	}
	return config

}

func Reload(config *models.Config) (reloaded bool) {
	configFile := locateConfig()
	if checkConfig(configFile) {
		newconfig := readConfig(configFile)
		*config = newconfig
		reloaded = true
	} else {
		reloaded = false
	}
	return reloaded
}

func locateConfig() (configFile string) {
	file := "jane.json"

	zero := *configParam()
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

func configParam() (configFile *string) {
	configFile = flag.String("config", "", "Location of the config file")
	flag.Parse()
	return configFile
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
