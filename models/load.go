package models

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
)

func Load() (config Config) {

	configFile := locateConfig()
	if checkConfig(configFile) {
		config = readConfig(configFile)
	} else {
		os.Exit(1)
	}
	return config

}

func Reload(config *Config) (reloaded bool) {
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
	if fileExists(zero) {
		return zero
	}

	first, _ := osext.ExecutableFolder()
	first += "/" + file
	if fileExists(first) {
		return first
	}

	second, _ := homedir.Dir()
	second += "/" + file
	if fileExists(second) {
		return second
	}

	third := "/etc/" + file
	if fileExists(third) {
		return third
	}

	return file

}

func configParam() (configFile *string) {
	configFile = flag.String("config", "", "Location of the config file")
	flag.Parse()
	return configFile
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readConfig(location string) (config Config) {
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
