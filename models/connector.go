package models

type Connector struct {
	Type         string
	ID           string
	Active       bool
	Server       string
	Login        string
	Pass         string
	From         string
	Key          string
	Image        string
	SuccessMatch string
	WarningMatch string
	FailureMatch string
	Checks       []struct {
		Name  string
		Check string
	}
	Commands []struct {
		Match  string
		Output string
		Cmd    string
		Args   string
	}
	Routes []Route
	Debug  bool
}
