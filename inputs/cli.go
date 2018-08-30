package inputs

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mmcquillan/hex/models"
	"github.com/mmcquillan/hex/parse"
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
		input, debug := parse.Flag(req, "--debug")
		if strings.TrimSpace(input) != "" {
			message := models.NewMessage()
			message.Attributes["hex.botname"] = config.BotName
			message.Attributes["hex.service"] = "cli"
			message.Attributes["hex.hostname"] = hostname
			message.Attributes["hex.user"] = user.Username
			message.Attributes["hex.input"] = input
			message.Debug = debug
			config.Logger.Debug("Cli Input - ID:" + message.Attributes["hex.id"])
			config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
			inputMsgs <- message
		} else {
			fmt.Print("\n> ")
		}
	}
}
