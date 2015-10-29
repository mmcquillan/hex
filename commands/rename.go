package commands

import (
	"github.com/mmcquillan/jane/configs"
	"strings"
)

func Rename(config *configs.Config, msg string) (results string) {
	name := strings.TrimSpace(msg)
	if strings.Contains(name, " ") {
		results = "Oops, can't do a name with a space."
	} else {
		config.JaneName = name
		results = "You may now call me '" + name + "'."
	}
	return results
}
