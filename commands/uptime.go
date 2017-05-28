package commands

import (
	"fmt"
	"time"

	"github.com/hexbotio/hex/models"
)

type Uptime struct {
}

func (x Uptime) Act(message *models.Message, states map[string]models.State, config *models.Config) {
	uptime := time.Now().Unix() - config.StartTime
	message.Response = append(message.Response, fmt.Sprintf("Up %d seconds.", uptime))
}
