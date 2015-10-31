package commands

import (
	"github.com/mmcquillan/jane/configs"
	"math/rand"
	"strings"
)

func Response(config *configs.Config, cmd string, msg string) (results string) {

	var match []string
	for _, r := range config.Responses {
		if strings.ToLower(r.In) == strings.ToLower(cmd) {
			match = append(match, r.Out)
		}
	}

	if len(match) > 0 {
		results = match[rand.Intn(len(match))]
		results = strings.Replace(results, "%msg%", msg, -1)
	} else {
		results = "Sorry, I do not understand what that means."
	}

	return results

}
