package listeners

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"log"
)

func Recovery(config *models.Config, listener models.Listener) {
	msg := config.Name + " Panic - " + listener.Name + " " + listener.Type + " Listener"
	if r := recover(); r != nil {
		log.Print(msg, r)
	}
	for _, d := range listener.Destinations {
		m := models.Message{
			Relays:      d.Relays,
			Target:      d.Target,
			Request:     "",
			Title:       msg,
			Description: "Check the log for more information and restart me to re-enable this listener.",
			Link:        "",
			Status:      "FAIL",
		}
		commands.Parse(config, m)
	}
}
