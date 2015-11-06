package models

type Config struct {
	Name      string
	LogFile   string
	Debug     bool
	Relays    []Relay
	Listeners []Listener
	Commands  []Command
}
