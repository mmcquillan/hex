package models

type Config struct {
	Name       string
	LogFile    string
	Debug      bool
	Connectors []Connector
	Commands   []Command
}
