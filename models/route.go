package models

type Route struct {
	Match struct {
		ConnectorType string
		ConnectorID   string
		Tags          string
		Target        string
		User          string
		Message       string
	}
	Publish struct {
		ConnectorID string
		Target      string
	}
}
