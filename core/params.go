package core

import (
	"flag"

	"github.com/hexbotio/hex/models"
)

func Params() (params models.Params) {
	configFile := flag.String("config", "", "Location of the config file")
	validate := flag.Bool("validate", false, "Validate the config file")
	flag.Parse()
	params.ConfigFile = *configFile
	params.Validate = *validate
	return params
}
