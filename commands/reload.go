package commands

import (
	"github.com/mmcquillan/jane/configs"
)

func Reload(config *configs.Config) (results string) {
	results = "Reloading configuration."
	*config = configs.Load()
	return results
}
