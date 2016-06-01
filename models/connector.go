package models

type Connector struct {
	Type     string
	ID       string
	Active   bool
	Server   string
	Port     string
	Login    string
	Pass     string
	Interval int
	File     string
	From     string
	Key      string
	Image    string
	Users    string
	Checks   []struct {
		Name   string
		Check  string
		Args   string
		Green  string
		Yellow string
		Red    string
	}
	Commands []struct {
		Match   string
		Output  string
		Outputs []string
		Cmd     string
		Args    string
		Help    string
	}
	Routes []Route
	Debug  bool
}
