package listeners

import (
	"bufio"
	"fmt"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/models"
	"os"
)

func Cli(config *models.Config, listener models.Listener) {
	fmt.Println("Starting in cli mode...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m := models.Message{
			Relays:      listener.Relays,
			Target:      listener.Target,
			Request:     scanner.Text(),
			Title:       "",
			Description: "",
			Link:        "",
			Status:      "",
		}
		commands.Parse(config, m)
	}
}
