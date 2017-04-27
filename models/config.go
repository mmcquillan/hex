package models

// Config Represents the values for configuring a jane server
type Config struct {
	BotName   string
	StartTime int64
	Debug     bool
	Version   string
	LogFile   string
	Aliases   []Alias
	Services  []Service
	Pipelines []Pipeline
}
