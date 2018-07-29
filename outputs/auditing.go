package outputs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mmcquillan/hex/models"
)

type Auditing struct {
}

func (x Auditing) Write(message models.Message, config models.Config) {
	who := " "
	if message.Attributes["hex.user"] != "" {
		who = who + message.Attributes["hex.user"] + " "
	}
	if message.Attributes["hex.channel"] != "" {
		if strings.HasPrefix(message.Attributes["hex.channel"], "#") {
			who = who + message.Attributes["hex.channel"] + " "
		} else {
			who = who + "DM "
		}
	}
	if message.Attributes["hex.schedule"] != "" {
		who = who + message.Attributes["hex.schedule"] + " "
	}
	if message.Attributes["hex.ipaddress"] != "" {
		who = who + message.Attributes["hex.ipaddress"] + " "
	}
	if message.Attributes["hex.url"] != "" {
		who = who + message.Attributes["hex.url"] + " "
	}
	out := fmt.Sprintf("%s [%s] %s >> %s\n", time.Now().Format(time.RFC3339), who, message.Attributes["hex.input"], message.Attributes["hex.rule.name"])
	file, err := os.OpenFile(config.AuditingFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		config.Logger.Error("Auditing File" + " - " + err.Error())
	}
	defer file.Close()
	if _, err = file.WriteString(out); err != nil {
		config.Logger.Error("Writing Audit File" + " - " + err.Error())
	}
}
