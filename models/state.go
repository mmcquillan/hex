package models

// State Struct
type State struct {
	LastRun  int64
	RunCount int64
	Success  bool
	Running  bool
}
