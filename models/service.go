package models

// Service Struct
type Service struct {
	Type   string
	ID     string
	Tags   string
	Active bool
	Config map[string]string
}
