package models

type Config struct {
	Name    string
	LogFile string
	Aliases []struct {
		Match  string
		Output string
	}
	Connectors []Connector
}
