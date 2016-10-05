package models

type Message struct {
	In struct {
		ConnectorType string
		ConnectorID   string
		Tags          string
		Target        string
		User          string
		Text          string
		Process       bool
		Timestamp     int64
	}
	Out struct {
		Target string
		Text   string
		Detail string
		Link   string
		Status string
	}
}
