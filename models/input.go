package models

// Input Struct
type Input struct {
	Type   string `json:"Type"`
	Name   string `json:"Name"`
	Target string `json:"Target"`
	Match  string `json:"Match"`
	ACL    string `json:"ACL,omitempty"`
	Help   string `json:"Help"`
	Hide   bool
}
