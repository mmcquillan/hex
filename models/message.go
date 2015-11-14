package models

type Message struct {
	Routes      []Route
	Source      string
	Request     string
	Title       string
	Description string
	Link        string
	Status      string
}
