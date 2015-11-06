package models

import (
	"fmt"
	"log"
	"os"
)

func Logging(config *Config) {

	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Error - Cannot get to log file at " + config.LogFile)
		panic(err)
	}
	log.SetOutput(f)

}
