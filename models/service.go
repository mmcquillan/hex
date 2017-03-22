package models

// Service Struct
type Service struct {
	Type   string
	Name   string
	Tags   string
	Active bool
	Config map[string]string
}
