package models

type Message struct {
	In struct {
		ConnectorType string
		ConnectorID   string
		Target        string
		User          string
		Text          string
		Process       bool
	}
	Out struct {
		Text   string
		Detail string
		Link   string
		Status string
	}
}
