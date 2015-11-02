package configs

import (
	"flag"
)

func Flags(config *Config) {

	Name := flag.String("Name", config.Name, "Set the name of your jane bot")
	flag.Parse()

	config.Name = *Name

}
