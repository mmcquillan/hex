package core

import (
	"fmt"
	"github.com/projectjane/jane/models"
	"log"
	"os"
)

func Logging(config *models.Config) {
	if config.LogFile == "" {
		log.SetOutput(os.Stdout)
	} else {
		if !FileExists(config.LogFile) {
			nf, err := os.Create(config.LogFile)
			if err != nil {
				fmt.Println("Error - Cannot create log file at " + config.LogFile)
			}
			nf.Close()
		}
		f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			fmt.Println("Error - Cannot get to log file at " + config.LogFile)
			panic(err)
		}
		log.SetOutput(f)
	}
	log.Print("---")
	log.Print("Starting jane bot...")
}
