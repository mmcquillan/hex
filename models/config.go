package models

// Config Represents the values for configuring a jane server
type Config struct {
	BotName    string
	Version    string
	LogFile    string
	Aliases    []Alias
	Services   []Service
	Connectors []Connector
	Routes     []Route
}
