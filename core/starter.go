package core

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/hexbotio/hex/models"
)

func Starter(config *models.Config) {

	// catch exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Received signal %v", sig)
			log.Print("Stopping hex bot...")
			os.Exit(0)
		}
	}()

	// setup logging
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

	// init logging feedback
	log.Print("")
	log.Print("Starting HexBot " + config.Version)
	log.Print("http://hexbot.io")

}
