package internals

import (
	"fmt"
	"time"

	"github.com/projectjane/jane/models"
)

type Uptime struct {
}

func (x Uptime) Act(message *models.Message, config *models.Config) {
	uptime := time.Now().Unix() - config.StartTime
	message.Response = append(message.Response, fmt.Sprintf("Up %d seconds.", uptime))
}
