package relays

import (
	"github.com/mmcquillan/jane/models"
	"strings"
)

func Output(config *models.Config, message models.Message) {
	for _, relay := range config.Relays {
		if relay.Active {
			if SendToRelay(relay.Type, message.Relays) {
				switch relay.Type {
				case "cli":
					Cli(config, relay, message)
				case "slack":
					Slack(config, relay, message)
				}
			}
		}
	}
}

func SendToRelay(rtype string, relays string) (send bool) {
	send = false
	if relays == "*" {
		send = true
	}
	r := strings.Split(relays, ",")
	for _, v := range r {
		if v == rtype {
			send = true
		}
	}
	return send
}
