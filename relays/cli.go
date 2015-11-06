package relays

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mmcquillan/jane/models"
)

func Cli(config *models.Config, relay models.Relay, message models.Message) {
	fmt.Println("")
	switch message.Status {
	case "SUCCESS":
		color.Set(color.FgGreen)
	case "WARN":
		color.Set(color.FgYellow)
	case "FAIL":
		color.Set(color.FgRed)
	}
	fmt.Println(message.Title)
	color.Unset()
	if message.Description != "" {
		fmt.Println(message.Description)
	}
}
