package models

type Message struct {
	Routes      []Route
	Target      string
	Request     string
	Title       string
	Description string
	Link        string
	Status      string
}
