package models

// Action Struct
type Action struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Command   string `json:"Command"`
	Success   string `json:"Success"`
	Failure   string `json:"Failure"`
	RunOnFail bool   `json:"RunOnFail"`
	Service   string `json:"Service"`
}
