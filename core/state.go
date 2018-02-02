package core

import (
	"sync"
	"time"

	"github.com/hexbotio/hex/models"
)

// State struct to control all things state
type State struct {
	updateLock *sync.Mutex
	States     map[string]models.State
}

// NewState function to initialize a state
func NewState(rules *map[string]models.Rule) (state *State) {
	state = &State{}
	state.updateLock = &sync.Mutex{}
	state.States = make(map[string]models.State)
	for _, rule := range *rules {
		state.States[rule.Id] = models.State{
			LastRun:  0,
			RunCount: 0,
			Success:  false,
			Running:  false,
		}
	}
	return state
}

// SetRunning function to mark if running or not
func (state State) SetRunning(ruleID string, running bool) {
	state.updateLock.Lock()
	defer state.updateLock.Unlock()
	s := state.States[ruleID]
	if running {
		s.Running = true
		s.RunCount = s.RunCount + 1
	} else {
		s.Running = false
		s.LastRun = time.Now().Unix()
	}
	state.States[ruleID] = s
}

// SetSuccess function to mark state success
func (state State) SetSuccess(ruleID string, success bool) {
	state.updateLock.Lock()
	defer state.updateLock.Unlock()
	s := state.States[ruleID]
	s.Success = success
	state.States[ruleID] = s
}
