package core

import (
	"log"
	"os"
	"os/signal"
)

func CatchExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Received signal %v", sig)
			log.Print("Stopping jane bot...")
			os.Exit(0)
		}
	}()
}
