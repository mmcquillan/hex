package connectors

import (
	"github.com/projectjane/jane/models"
	"log"
)

func Recovery(connector models.Connector) {
	msg := "Panic - " + connector.ID + " " + connector.Type + " Connector"
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
}
