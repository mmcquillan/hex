package inputs

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

// Command struct
type Command struct {
}

// Read function
func (x Command) Read(inputMsgs chan<- models.Message, config models.Config) {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	input, debug := parse.Flag(config.Command, "--debug")
	if strings.TrimSpace(input) != "" {
		message := models.NewMessage()
		message.Attributes["hex.service"] = "command"
		message.Attributes["hex.hostname"] = hostname
		message.Attributes["hex.user"] = user.Username
		message.Attributes["hex.input"] = input
		message.Debug = debug
		config.Logger.Debug("Command Input - ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
		inputMsgs <- message
	}
}
