package outputs

import (
	"github.com/mmcquillan/jane/configs"
)

func Output(config *configs.Config, message Message) {
	if config.Interactive {
		CLI(config, message)
	} else {
		Slack(config, message)
	}
}
