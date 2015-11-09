package connectors

import (
	"bufio"
	"fmt"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"log"
	"os"
	"strings"
)

type Cli struct {
}

func (x Cli) Run(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	fmt.Println("Starting in cli mode...\n")
	fmt.Print(config.Name + "> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		if req == "exit" {
			log.Print("Exiting jane bot by command line")
			os.Exit(0)
		}
		for _, r := range connector.Routes {
			if config.Debug {
				log.Print("Processing CLI")
			}
			if strings.Contains(req, r.Match) || r.Match == "*" {
				m := models.Message{
					Relays:      r.Relays,
					Target:      r.Target,
					Request:     req,
					Title:       "",
					Description: "",
					Link:        "",
					Status:      "",
				}
				commands.Parse(config, m)
			}
		}
		fmt.Print("\n" + config.Name + "> ")
	}
}
