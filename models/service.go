package models

// Service Struct
type Service struct {
	BotName string
	Type    string
	Name    string
	Active  bool
	Config  map[string]string
}
