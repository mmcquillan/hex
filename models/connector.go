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
	Routes []Route
}
