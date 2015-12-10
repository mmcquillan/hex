package models

type Message struct {
	Routes []Route
	In     struct {
		Source    string
		User      string
		Text      string
		Process   bool
		Reprocess bool
	}
	Out struct {
		Text   string
		Detail string
		Link   string
		Status string
	}
}
