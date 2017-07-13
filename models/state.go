package models

// State Struct
type State struct {
	LastRun     int64
	LastChange  int64
	LastAlert   int64
	RunStart    int64
	RunCount    int64
	LastRunTime int64
	AvgRunTime  int64
	Success     bool
	Running     bool
	Alert       bool
}
