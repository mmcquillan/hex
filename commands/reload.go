package commands

import (
	"github.com/mmcquillan/jane/models"
)

func Reload(command models.Command, config *models.Config) (results string) {
	results = command.Output
	*config = models.Load()
	return results
}
