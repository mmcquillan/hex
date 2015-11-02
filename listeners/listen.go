package listeners

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/relays"
	"strconv"
	"time"
)

func Listen(config *configs.Config) {

	// init
	now := time.Now()
	messages := make([]relays.Message, 0)
	deployMarker := strconv.FormatInt(now.Unix(), 10) + "000"

	// general loop
	for {

		// deploys
		deployMarker, messages = Deploys(config, deployMarker)
		for _, m := range messages {
			relays.OutputAll(config, m)
		}

		// wait
		time.Sleep(120 * time.Second)

	}

}
