package core

import (
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func LoadConfig(params models.Params, version string) (config models.Config) {
	configFile := locateConfig(params)
	if checkConfig(configFile) {
		config = readConfig(configFile)
		subConfig(&config)
		if params.Validate {
			fmt.Println("SUCCESS - Config file is valid: " + configFile)
			os.Exit(0)
		}
	} else {
		os.Exit(1)
	}
	config.Version = version
	return config
}

func locateConfig(params models.Params) (configFile string) {
	tryfile := ""
	file := "jane.json"

	// first try env param
	tryfile = os.Getenv("JANE_CONFIG")
	if FileExists(tryfile) {
		return tryfile
	}

	// second try param
	tryfile = params.ConfigFile
	if FileExists(tryfile) {
		return tryfile
	}

	// third try jane config in current executable dir
	tryfile, _ = osext.ExecutableFolder()
	tryfile += "/" + file
	if FileExists(tryfile) {
		return tryfile
	}

	// fourth try jane config in home dir
	tryfile, _ = homedir.Dir()
	tryfile += "/" + file
	if FileExists(tryfile) {
		return tryfile
	}

	// fifth try jane config in /etc
	tryfile = "/etc/" + file
	if FileExists(tryfile) {
		return tryfile
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

func subConfig(config *models.Config) {
	for i := 0; i < len(config.Connectors); i++ {
		config.Connectors[i].Server = parse.SubstituteInputs(config.Connectors[i].Server)
		config.Connectors[i].Port = parse.SubstituteInputs(config.Connectors[i].Port)
		config.Connectors[i].Login = parse.SubstituteInputs(config.Connectors[i].Login)
		config.Connectors[i].Pass = parse.SubstituteInputs(config.Connectors[i].Pass)
		config.Connectors[i].Key = parse.SubstituteInputs(config.Connectors[i].Key)
		if os.Getenv("JANE_DEBUG") != "" {
			if strings.ToLower(os.Getenv("JANE_DEBUG")) == "true" {
				config.Connectors[i].Debug = true
			} else {
				config.Connectors[i].Debug = false
			}
		}
	}
	for i := 0; i < len(config.Routes); i++ {
		if config.Routes[i].Match.ConnectorType == "" {
			config.Routes[i].Match.ConnectorType = "*"
		}
		if config.Routes[i].Match.ConnectorID == "" {
			config.Routes[i].Match.ConnectorID = "*"
		}
		if config.Routes[i].Match.Tags == "" {
			config.Routes[i].Match.Tags = "*"
		}
		if config.Routes[i].Match.Target == "" {
			config.Routes[i].Match.Target = "*"
		}
		if config.Routes[i].Match.User == "" {
			config.Routes[i].Match.User = "*"
		}
		if config.Routes[i].Match.Message == "" {
			config.Routes[i].Match.Message = "*"
		}
	}
	if os.Getenv("JANE_LOGFILE") != "" {
		config.LogFile = os.Getenv("JANE_LOGFILE")
	}
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
