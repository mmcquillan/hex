package configs

type Config struct {
	LogFile        string
	Interactive    bool
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
