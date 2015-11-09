package connectors

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"log"
)

func Recovery(config *models.Config, connector models.Connector) {
	msg := config.Name + " Panic - " + connector.Name + " " + connector.Type + " Connector"
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
	for _, r := range connector.Routes {
		m := models.Message{
			Relays:      r.Relays,
			Target:      r.Target,
			Request:     "",
			Title:       msg,
			Description: "Check the log for more information and restart me to re-enable this connector.",
			Link:        "",
			Status:      "FAIL",
		}
		commands.Parse(config, m)
	}
}
