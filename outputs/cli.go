package outputs

import (
	"fmt"
	"github.com/mmcquillan/jane/configs"
)

func CLI(config *configs.Config, message Message) {
	fmt.Println("Name:   " + config.Name)
	fmt.Println("Title:  " + message.Title)
	fmt.Println("Desc:   " + message.Description)
	fmt.Println("Link:   " + message.Link)
	fmt.Println("Status: " + message.Status)
	fmt.Println("")
}
