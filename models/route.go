package models

type Route struct {
	Matches []struct {
		ConnectorType string
		ConnectorID   string
		Target        string
		User          string
		Message       string
	}
	Connectors string
	Targets    string
}
