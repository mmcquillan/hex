package models

type Config struct {
	LogFile string
	Aliases []struct {
		Match  string
		Output string
	}
	Connectors []Connector
}
