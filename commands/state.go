package commands

import (
	"fmt"
	"time"

	"github.com/hexbotio/hex/models"
)

type State struct {
}

func (x State) Act(message *models.Message, states map[string]models.State, config *models.Config) {
	now := time.Now().Unix()
	for pipeline, state := range states {
		if state.Alert {
			r := fmt.Sprintf("%s: %t", pipeline, state.Success)
			if state.LastRun > 0 && !state.Running {
				r = fmt.Sprintf("%s for %d secs (checked %d sec ago)", r, now-state.LastChange, now-state.LastRun)
			}
			if state.LastRun > 0 && state.Running {
				r = fmt.Sprintf("%s for %d secs (running)", r, now-state.LastChange)
			}
			message.Response = append(message.Response, r)
		}
	}
}
