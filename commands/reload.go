package commands

import (
	"github.com/projectjane/jane/models"
)

func Reload(command models.Command, config *models.Config) (results string) {
	results = command.Output
	if models.Reload(config) {
		results = command.Output
	} else {
		results = "Configuration is invalid, please check it."
	}
	return results
}
