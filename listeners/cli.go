package listeners

import (
	"bufio"
	"fmt"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"os"
	"strings"
)

func Cli(config *models.Config, listener models.Listener) {
	fmt.Println("Starting in cli mode...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, d := range listener.Destinations {
			req := scanner.Text()
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
