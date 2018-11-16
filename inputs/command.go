package inputs

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mmcquillan/hex/models"
)

// Command struct
type Command struct {
}

// Read function
func (x Command) Read(inputMsgs chan<- models.Message, config models.Config) {
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	if strings.TrimSpace(config.Command) != "" {
		message := models.NewMessage()
		message.Attributes["hex.service"] = "command"
		message.Attributes["hex.hostname"] = hostname
		message.Attributes["hex.user"] = user.Username
		message.Attributes["hex.input"] = config.Command
		config.Logger.Debug("Command Input - ID:" + message.Attributes["hex.id"])
		config.Logger.Trace(fmt.Sprintf("Message: %+v", message))
		inputMsgs <- message
	}
}
