package models

// State Struct
type State struct {
	LastRun    int64
	LastChange int64
	LastAlert  int64
	Success    bool
	Running    bool
	Alert      bool
}
