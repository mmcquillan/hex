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
			LastRun:     0,
			LastChange:  time.Now().Unix(),
			LastAlert:   0,
			RunStart:    0,
			RunCount:    0,
			LastRunTime: 0,
			AvgRunTime:  0,
			Success:     true,
			Running:     false,
			Alert:       p.Alert,
		}
	}
	return state
}

func (state State) SetRunning(pipeline string, running bool) {
	s := state.States[pipeline]
	s.Running = running
	if running {
		s.RunStart = time.Now().Unix()
		s.RunCount = s.RunCount + 1
	} else {
		runEnd := time.Now().Unix()
		s.LastRunTime = runEnd - s.RunStart
		s.AvgRunTime = ((s.AvgRunTime * (s.RunCount - 1)) + s.LastRunTime) / s.RunCount
		s.RunStart = 0
	}
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
