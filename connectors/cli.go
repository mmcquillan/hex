package connectors

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"log"
	"os"
	"os/user"
)

type Cli struct {
}

func (x Cli) Listen(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	u, err := user.Current()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Starting in cli mode...\n")
	fmt.Print(config.Name + "> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		if req == "exit" {
			log.Print("Exiting jane bot by command line")
			os.Exit(0)
		}
		if connector.Debug {
			log.Print("Logging msg: " + req)
		}
		m := models.Message{
			Routes:      connector.Routes,
			Source:      u.Username,
			Request:     req,
			Title:       "",
			Description: "",
			Link:        "",
			Status:      "",
		}
		commands.Parse(config, &m)
		Broadcast(config, m)
	}
}

func (x Cli) Command(config *models.Config, message *models.Message) {
	return
}

func (x Cli) Publish(config *models.Config, connector models.Connector, message models.Message, target string) {
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
	fmt.Print("\n" + config.Name + "> ")
}
