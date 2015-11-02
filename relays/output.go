package relays

import (
	"github.com/mmcquillan/jane/configs"
)

func Output(config *configs.Config, relay configs.Relay, message Message) {
	switch relay.Type {
	case "cli":
		CliOut(config, relay, message)
	case "slack":
		SlackOut(config, relay, message)
	}
}

func OutputAll(config *configs.Config, message Message) {
	for _, relay := range config.Relays {
		if relay.Active {
			Output(config, relay, message)
		}
	}
}
