package models

// Pipeline Struct
type Pipeline struct {
	Name    string `json:"Name"`
	Active  bool   `json:"Active"`
	Track   bool   `json:"Track"`
	Inputs  []Input
	Actions []Action
	Outputs []Output
}
