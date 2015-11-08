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

func Cli(config *models.Config, connector models.Connector) {
	defer Recovery(config, connector)
	fmt.Println("Starting in cli mode...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, d := range connector.Destinations {
			req := scanner.Text()
			if config.Debug {
				log.Print("Processing CLI")
			}
			if strings.Contains(req, d.Match) || d.Match == "*" {
				m := models.Message{
					Relays:      d.Relays,
					Target:      d.Target,
					Request:     req,
					Title:       "",
					Description: "",
					Link:        "",
					Status:      "",
				}
				commands.Parse(config, m)
			}
		}
	}
}
