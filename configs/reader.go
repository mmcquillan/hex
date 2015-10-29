package configs

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

func ReadConfig(location string) (config Config) {

	if _, err := toml.DecodeFile(location, &config); err != nil {
		log.Println(err)
	}
	return config

}

func CheckConfig(location string) (exists bool) {
	exists = true
	if _, err := os.Stat(location); os.IsNotExist(err) {
		exists = false
	}
	return exists
}
