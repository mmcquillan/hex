package models

type Message struct {
	Routes []Route
	In     struct {
		Source  string
		User    string
		Text    string
		Process bool
	}
	Out struct {
		Text   string
		Detail string
		Link   string
		Status string
	}
}
