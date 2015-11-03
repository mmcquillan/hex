package models

type Listener struct {
	Type         string
	Name         string
	Resource     string
	Target       string
	Relays       string
	SuccessMatch string
	FailureMatch string
	Active       bool
}
