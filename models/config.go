package models

type Config struct {
	Name       string
	LogFile    string
	Debug      bool
	Relays     []Relay
	Connectors []Connector
	Commands   []Command
}
