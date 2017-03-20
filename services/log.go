package services

import (
	"fmt"
	"github.com/projectjane/jane/models"
	"log"
	"os"
	"time"
)

type Log struct {
}

func (x Log) Input(inputMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Log) Command(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Log) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	file, err := os.OpenFile(connector.File, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Print(err)

	}
	defer file.Close()
	for {
		message := <-outputMsgs
		if _, err = file.WriteString(time.Now().Format(time.RFC3339) + " " + fmt.Sprintf("%+v", message) + "\n"); err != nil {
			log.Print(err)

		}
	}
}

func (x Log) Help(connector models.Connector) (help []string) {
	return
}
