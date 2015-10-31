package configs

type Config struct {
	Name        string
	LogFile     string
	Interactive bool
	Debug       bool

	// temp config start
	JaneEmoji      string
	JaneChannel    string
	SlackToken     string
	BambooUrl      string
	BambooUser     string
	BambooPass     string
	BambooChannels map[string]string
	// temp config end

	Listeners []struct {
		Type         string // rss,slack,
		Input        string // the connection method
		Output       string // slack,cli
		Filter       string // string filter
		SuccessMatch string // string to indicate success
		FailureMatch string // string to indicate failure
		Active       bool
	}
	Responses []struct {
		In  string
		Out string
	}
}
