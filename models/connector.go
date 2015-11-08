package models

type Connector struct {
	Type         string
	Name         string
	Active       bool
	Server       string
	Login        string
	Pass         string
	Key          string
	SuccessMatch string
	WarningMatch string
	FailureMatch string
	Checks       []struct {
		Name  string
		Check string
	}
	Destinations []struct {
		Match  string
		Relays string
		Target string
	}
}
