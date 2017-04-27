package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

func Config(params models.Params, version string) (config models.Config) {
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
	config.StartTime = time.Now().Unix()
	return config
}

func locateConfig(params models.Params) (configFile string) {
	tryfile := ""
	file := "hex.json"

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

	// third try hex config in current executable dir
	tryfile, _ = osext.ExecutableFolder()
	tryfile += "/" + file
	if FileExists(tryfile) {
		return tryfile
	}

	// fourth try hex config in home dir
	tryfile, _ = homedir.Dir()
	tryfile += "/" + file
	if FileExists(tryfile) {
		return tryfile
	}

	// fifth try hex config in /etc
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
	if os.Getenv("JANE_BOT_NAME") != "" {
		config.BotName = os.Getenv("JANE_BOT_NAME")
	} else if config.BotName == "" {
		config.BotName = "hex"
	}
	if os.Getenv("JANE_DEBUG") != "" {
		if strings.ToLower(os.Getenv("JANE_DEBUG")) == "true" || config.Debug {
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
