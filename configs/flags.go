package configs

import (
	"flag"
)

func Flags(config *Config) {

	Interactive := flag.Bool("Interactive", config.Interactive, "Run jane in interactive cli mode")
	JaneName := flag.String("JaneName", config.JaneName, "Set the name of your jane bot")
	flag.Parse()

	config.Interactive = *Interactive
	config.JaneName = *JaneName

}
