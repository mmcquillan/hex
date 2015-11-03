package relays

import (
	"fmt"
	"github.com/mmcquillan/jane/models"
)

func Cli(config *models.Config, relay models.Relay, message models.Message) {
	fmt.Println(config.Name + ": " + message.Title)
	if message.Description != "" {
		fmt.Println(message.Description)
	}
	fmt.Println("")
}
