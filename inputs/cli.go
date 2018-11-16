package inputs

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mmcquillan/hex/models"
)

// Cli struct
type Cli struct {
}

// Read function
func (x Cli) Read(inputMsgs chan<- models.Message, config models.Config) {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	fmt.Print("Starting in cli mode...\n")
	fmt.Print("\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		if strings.TrimSpace(req) != "" {
			message := models.NewMessage()
			message.Attributes["hex.botname"] = config.BotName
			message.Attributes["hex.service"] = "cli"
			message.Attributes["hex.hostname"] = hostname
			message.Attributes["hex.user"] = user.Username
			message.Attributes["hex.input"] = req
			config.Logger.Debug("Cli Input - ID:" + message.Attributes["hex.id"])
			config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
			inputMsgs <- message
		} else {
			fmt.Print("\n> ")
		}
	}
}
