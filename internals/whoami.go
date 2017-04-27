package internals

import (
	"github.com/projectjane/jane/models"
)

type Whoami struct {
}

func (x Whoami) Act(message *models.Message, config *models.Config) {
	message.Response = append(message.Response, message.Inputs["jane.user"])
}
