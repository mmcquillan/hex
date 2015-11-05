package models

type Config struct {
	Name      string
	LogFile   string
	Debug     bool
	NewRelic  string
	Relays    []Relay
	Listeners []Listener
	Commands  []Command
}
