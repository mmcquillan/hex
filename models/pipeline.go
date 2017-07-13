package models

// Pipeline Struct
type Pipeline struct {
	Name    string `json:"Name"`
	Active  bool   `json:"Active"`
	Alert   bool   `json:"Alert"`
	Remind  int64  `json:"Remind"`
	Inputs  []Input
	Actions []Action
	Outputs []Output
}
