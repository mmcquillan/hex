package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

func Config(config *models.Config) {
	locateConfigFile(config)
	if checkConfig(config) {
		readConfig(config)
		subConfig(config)
		configRules(config)
		if config.Validate {
			fmt.Println("SUCCESS - Config file is valid: " + config.ConfigFile)
			os.Exit(0)
		}
	} else {
		os.Exit(1)
	}
	config.StartTime = time.Now().Unix()
}

func locateConfigFile(config *models.Config) {
	locations := []string{
		os.Getenv("HEX_CONFIG"),
		config.ConfigFile,
		"/etc/hex.json",
		"/etc/hex/hex.json",
	}
	for _, location := range locations {
		if FileExists(location) {
			config.ConfigFile = location
			return
		}
	}
}

func readConfig(config *models.Config) {
	file, err := ioutil.ReadFile(config.ConfigFile)
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Print(err)
	}
}

func subConfig(config *models.Config) {

	// handle bot name
	if os.Getenv("HEX_BOT_NAME") != "" {
		config.BotName = os.Getenv("HEX_BOT_NAME")
	} else if config.BotName == "" {
		config.BotName = "hex"
	}
	if strings.HasPrefix(config.BotName, "@") {
		config.BotName = strings.Replace(config.BotName, "@", "", -1)
	}

	// handle workspace
	if os.Getenv("HEX_WORKSPACE") != "" {
		config.BotName = os.Getenv("HEX_WORKSPACE")
	} else if config.Workspace == "" {
		config.Workspace = "/tmp"
	}
	if !strings.HasSuffix(config.Workspace, "/") {
		config.Workspace = config.Workspace + "/"
	}
	if _, err := os.Stat(config.Workspace); os.IsNotExist(err) {
		fmt.Println("ERROR - Workspace directory is invalid: " + config.Workspace)
		os.Exit(1)
	}

	// handle debug
	if os.Getenv("HEX_DEBUG") != "" {
		if strings.ToLower(os.Getenv("HEX_DEBUG")) == "true" || config.Debug {
			config.Debug = true
		} else {
			config.Debug = false
		}
	}
	for i := 0; i < len(config.Services); i++ {
		for k, v := range config.Services[i].Config {
			config.Services[i].Config[k] = parse.SubstituteEnv(v)
		}
		config.Services[i].BotName = config.BotName
	}

	// handle logfile
	if os.Getenv("HEX_LOGFILE") != "" {
		config.LogFile = os.Getenv("HEX_LOGFILE")
	}

}

func configRules(config *models.Config) {

	// check for service name uniqueness
	serviceChk := make(map[string]bool)
	for _, service := range config.Services {
		serviceChk[service.Name] = true
	}
	if len(config.Services) > len(serviceChk) {
		log.Print("ERROR - Duplicate Service Names exist")
		os.Exit(1)
	}

	// check for pipeline name uniqueness
	pipelineChk := make(map[string]bool)
	for _, pipeline := range config.Pipelines {
		pipelineChk[pipeline.Name] = true
	}
	if len(config.Pipelines) > len(pipelineChk) {
		log.Print("ERROR - Duplicate Pipeline Names exist")
		os.Exit(1)
	}

}

func checkConfig(config *models.Config) (exists bool) {
	exists = true
	if _, err := os.Stat(config.ConfigFile); os.IsNotExist(err) {
		fmt.Println("Error - Missing a config file")
		exists = false
	}
	if exists {
		file, err := ioutil.ReadFile(config.ConfigFile)
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

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
