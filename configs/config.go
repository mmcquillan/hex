package configs

type Config struct {
	JaneName       string
	JaneEmoji      string
	JaneChannel    string
	SlackToken     string
	BambooUser     string
	BambooPass     string
	BambooChannels map[string]string
}
