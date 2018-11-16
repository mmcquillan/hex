package outputs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mmcquillan/hex/models"
)

type Auditing struct {
	AuditFile *os.File
}

func (x Auditing) Write(message models.Message, config models.Config) {

	// create file if specified
	if config.AuditingFile != "" && x.AuditFile == nil {
		if validateLog(config) {
			file, err := os.OpenFile(config.AuditingFile, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				config.Logger.Error("Auditing File" + " - " + err.Error())
			}
			x.AuditFile = file
			defer x.AuditFile.Close()
		} else {
			return
		}
	}

	// set who
	who := ""
	if message.Attributes["hex.user"] != "" {
		who = who + message.Attributes["hex.user"]
	}
	if message.Attributes["hex.channel"] != "" {
		if strings.HasPrefix(message.Attributes["hex.channel"], "#") {
			who = who + message.Attributes["hex.channel"]
		} else {
			who = who + "DM"
		}
	}
	if message.Attributes["hex.schedule"] != "" {
		who = who + message.Attributes["hex.schedule"]
	}
	if message.Attributes["hex.ipaddress"] != "" {
		who = who + message.Attributes["hex.ipaddress"]
	}
	if message.Attributes["hex.url"] != "" {
		who = who + message.Attributes["hex.url"]
	}

	// log out
	out := fmt.Sprintf("%s [AUDIT] %s: %s >> %s\n", time.Now().Format(time.RFC3339), who, message.Attributes["hex.input"], message.Attributes["hex.rule.name"])
	if config.AuditingFile != "" {
		if _, err := x.AuditFile.WriteString(out); err != nil {
			config.Logger.Error("Writing Audit File" + " - " + err.Error())
		}
	} else {
		fmt.Print(out)
	}

}

func validateLog(config models.Config) bool {
	if _, err := os.Stat(config.AuditingFile); os.IsNotExist(err) {
		nf, err := os.Create(config.AuditingFile)
		if err != nil {
			config.Logger.Error("ERROR: Cannot create Auditing File at " + config.AuditingFile + " - " + err.Error())
			return false
		}
		nf.Close()
		config.Logger.Info("Created Auditing File " + config.AuditingFile)
	}
	return true
}
