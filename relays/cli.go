package relays

import (
	"bufio"
	"fmt"
	"github.com/mmcquillan/jane/configs"
	"os"
)

func CliIn(config *configs.Config, relay configs.Relay) {
	fmt.Println("Starting in cli mode...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("")
		Parse(config, relay, "", scanner.Text())
	}
}

func CliOut(config *configs.Config, relay configs.Relay, message Message) {
	fmt.Println("Name:   " + config.Name)
	fmt.Println("Image:  " + relay.Image)
	fmt.Println("Title:  " + message.Title)
	fmt.Println("Desc:   " + message.Description)
	fmt.Println("Link:   " + message.Link)
	fmt.Println("Status: " + message.Status)
	fmt.Println("")
}
