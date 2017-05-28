package core

import (
	"time"

	"github.com/hexbotio/hex/models"
)

type State struct {
	States map[string]models.State
}

func NewState(config *models.Config) (state *State) {
	state = &State{}
	state.States = make(map[string]models.State)
	for _, p := range config.Pipelines {
		state.States[p.Name] = models.State{
			LastRun:    0,
			LastChange: 0,
			LastAlert:  0,
			Success:    true,
			Running:    false,
		}
	}
	return state
}

func (state State) SetRunning(pipeline string, running bool) {
	s := state.States[pipeline]
	s.Running = running
	state.States[pipeline] = s
}

func (state State) SetState(pipeline string, success bool) {
	s := state.States[pipeline]
	if s.Success != success {
		s.LastChange = time.Now().Unix()
	}
	s.Success = success
	state.States[pipeline] = s
}

func (state State) SetLastRun(pipeline string) {
	s := state.States[pipeline]
	s.LastRun = time.Now().Unix()
	state.States[pipeline] = s
}

func (state State) SetLastAlert(pipeline string) {
	s := state.States[pipeline]
	s.LastAlert = time.Now().Unix()
	state.States[pipeline] = s
}
