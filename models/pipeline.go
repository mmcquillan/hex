package models

// Pipeline Struct
type Pipeline struct {
	Name    string `json:"Name"`
	Active  bool   `json:"Active"`
	Alert   bool   `json:"Alert"`
	Inputs  []Input
	Actions []Action
	Outputs []Output
}
