package listeners

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/outputs"
	"strconv"
	"time"
)

func Listen(config *configs.Config) {

	// init
	now := time.Now()
	messages := make([]outputs.Message, 0)
	deployMarker := strconv.FormatInt(now.Unix(), 10) + "000"

	// general loop
	for {

		// deploys
		deployMarker, messages = Deploys(config, deployMarker)
		for _, m := range messages {
			outputs.Output(config, m)
		}

		// wait
		time.Sleep(120 * time.Second)

	}

}
