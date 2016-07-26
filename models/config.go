package models

type Config struct {
	Version string
	LogFile string
	Aliases []struct {
		Match  string
		Output string
	}
	Connectors []Connector
}
