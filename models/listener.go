package models

type Listener struct {
	Type         string
	Name         string
	Resource     string
	SuccessMatch string
	FailureMatch string
	Active       bool
	Destinations []struct {
		Match  string
		Relays string
		Target string
	}
}
