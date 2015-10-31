package configs

import (
	"flag"
)

func Flags(config *Config) {

	Interactive := flag.Bool("Interactive", config.Interactive, "Run jane in interactive cli mode")
	Name := flag.String("Name", config.Name, "Set the name of your jane bot")
	flag.Parse()

	config.Interactive = *Interactive
	config.Name = *Name

}
