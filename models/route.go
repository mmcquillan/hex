package models

type Route struct {
	Matches []struct {
		ConnectorType string
		ConnectorID   string
		Tags          string
		Target        string
		User          string
		Message       string
	}
	Connectors string
	Targets    string
}
