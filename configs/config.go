package configs

type Config struct {
	LogFile        string
	Debug          bool
	JaneName       string
	JaneEmoji      string
	JaneChannel    string
	SlackToken     string
	BambooUrl      string
	BambooUser     string
	BambooPass     string
	BambooChannels map[string]string
}
