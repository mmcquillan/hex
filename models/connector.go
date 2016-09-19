package models

type Connector struct {
	Type     string
	ID       string
	Tags     string
	Active   bool
	Server   string
	Port     string
	Login    string
	Pass     string
	File     string
	From     string
	Key      string
	Image    string
	Users    string
	Commands []Command
	Debug    bool
}
