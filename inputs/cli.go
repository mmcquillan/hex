package inputs

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

// Cli struct
type Cli struct {
}

// Read function
func (x Cli) Read(inputMsgs chan<- models.Message, config models.Config) {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	fmt.Println("Starting in cli mode...\n")
	fmt.Print("> ", config.BotName, " ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		fmt.Print("\n> ", config.BotName, " ")
		input, debug := parse.Flag(req, "--debug")
		if strings.TrimSpace(input) != "" {
			message := models.NewMessage()
			message.Attributes["hex.service"] = "cli"
			message.Attributes["hex.hostname"] = hostname
			message.Attributes["hex.user"] = user.Username
			message.Attributes["hex.input"] = input
			message.Debug = debug
			inputMsgs <- message
		}
	}
}
