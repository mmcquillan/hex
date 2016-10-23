package models

// Connector Struct representing a connector from the config
type Connector struct {
	Type              string
	ID                string
	Tags              string
	Active            bool
	Server            string
	Port              string
	Login             string
	Pass              string
	File              string
	From              string
	Key               string
	Secret            string
	AccessToken       string
	AccessTokenSecret string
	Image             string
	Users             string
	Commands          []Command
	Debug             bool
	BotName           string
	Filter            []string
}
