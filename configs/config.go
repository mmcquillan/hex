package configs

type Config struct {
	Name    string
	LogFile string
	Debug   bool

	// temp config start
	BambooUrl      string
	BambooUser     string
	BambooPass     string
	BambooChannels map[string]string
	// temp config end

	Relays    []Relay
	Listeners []Listener
	Responses []struct {
		In  string
		Out string
	}
}

type Relay struct {
	Type     string
	Image    string
	Resource string
	Active   bool
}

type Listener struct {
	Type         string // rss,slack,
	Name         string // a friendly name
	Resource     string // the connection method
	Output       string // slack,cli
	Filter       string // string filter
	SuccessMatch string // string to indicate success
	FailureMatch string // string to indicate failure
	Active       bool
}
