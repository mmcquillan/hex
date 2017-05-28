package models

// Config Represents the values for configuring a hex server
type Config struct {
	BotName    string
	StartTime  int64
	Debug      bool
	ConfigFile string
	Validate   bool
	Version    string
	LogFile    string
	Workspace  string
	Services   []Service
	Pipelines  []Pipeline
}
