package actions

import (
	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

type Format struct {
}

func (x Format) Act(action models.Action, message *models.Message, config *models.Config) {
	results := parse.Substitute(action.Command, message.Inputs)
	message.Response = append(message.Response, results)
	message.Success = ActionSuccess(results, action.Success, action.Failure)
}
