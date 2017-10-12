package commands

import (
	"fmt"

	"github.com/hexbotio/hex/models"
)

type Rules struct {
}

func (x Rules) Act(message *models.Message, rules *map[string]models.Rule, config models.Config) {
	res := "Rules:\n"
	for key, rule := range *rules {
		res = fmt.Sprintf("%s%s - %+v\n", res, key, rule)
	}
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "rules",
		Response: res,
	})
}
