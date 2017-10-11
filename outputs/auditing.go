package outputs

import (
	"fmt"
	"os"
	"time"

	"github.com/hexbotio/hex/models"
)

type Auditing struct {
}

func (x Auditing) Write(message models.Message, config models.Config) {
	file, err := os.OpenFile(config.AuditingFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		config.Logger.Error("Auditing File", err)
	}
	defer file.Close()
	out := fmt.Sprintf("%s %+v\n", time.Now().Format(time.RFC3339), message)
	if _, err = file.WriteString(out); err != nil {
		config.Logger.Error("Writing Audit File", err)
	}
}
