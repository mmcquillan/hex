package models

type Config struct {
	Name       string
	LogFile    string
	Connectors []Connector
	Commands   []Command
}
