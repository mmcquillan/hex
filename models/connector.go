package models

type Connector struct {
	Type   string
	ID     string
	Active bool
	Server string
	Port   string
	Login  string
	Pass   string
	From   string
	Key    string
	Image  string
	Users  string
	Checks []struct {
		Name  string
		Check string
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
