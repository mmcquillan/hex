package models

type Config struct {
	Version    string
	LogFile    string
	Aliases    []Alias
	Connectors []Connector
}
