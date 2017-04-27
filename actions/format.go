package actions

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

type Format struct {
}

func (x Format) Act(action models.Action, message *models.Message, config *models.Config) {
	results := parse.Substitute(action.Command, message.Inputs)
	message.Response = append(message.Response, results)
	message.Success = ActionSuccess(results, action.Success, action.Failure)
}
