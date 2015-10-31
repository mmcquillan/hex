package outputs

import (
	"fmt"
	"github.com/mmcquillan/jane/configs"
)

func CLI(config *configs.Config, message Message) {
	fmt.Println(config.Name + ": " + message.Title)
	if message.Description != "" {
		fmt.Println(message.Description)
	}
	fmt.Println("")
}

func colorMe(status string) (color string) {
	switch status {
	case "Successful":
		color = "good"
	case "SUCCESS":
		color = "good"
	case "Failed":
		color = "danger"
	case "FAILED":
		color = "danger"
	default:
		color = "#DDDDDD"
	}
	return color
}
