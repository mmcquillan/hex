package internals

import (
	"github.com/projectjane/jane/models"
)

type Ping struct {
}

func (x Ping) Act(message *models.Message, config *models.Config) {
	message.Response = append(message.Response, "pong")
}
