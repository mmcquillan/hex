package configs

type Config struct {
	SlackerName    string
	SlackerEmoji   string
	SlackerChannel string
	SlackToken     string
	BambooUser     string
	BambooPass     string
	BambooChannels map[string]string
}
