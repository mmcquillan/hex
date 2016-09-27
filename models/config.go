package models

type Config struct {
	BotName    string
	Version    string
	LogFile    string
	Aliases    []Alias
	Connectors []Connector
	Routes     []Route
}
