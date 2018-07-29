package commands

import (
	"fmt"

	"github.com/mmcquillan/hex/models"
)

func Rules(message *models.Message, rules *map[string]models.Rule, config models.Config) {
	res := "Rules:\n\n"
	for key, rule := range *rules {
		res = fmt.Sprintf("%s%s - %+v\n\n", res, key, rule)
	}
	message.Attributes["hex.rule.format"] = "true"
	message.Outputs = append(message.Outputs, models.Output{
		Rule:     "rules",
		Response: res,
		Success:  true,
	})
}
