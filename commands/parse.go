package commands

import (
	"bitbucket.org/prysm/devops-robot/configs"
	"github.com/nlopes/slack"
	"strings"
)

func Parse(config *configs.Config, channel string, msg string) {

	// make sure they are talking to and not about us
	tokmsg := strings.Split(strings.TrimSpace(msg), " ")
	if strings.ToLower(tokmsg[0]) != strings.ToLower(config.SlackerName) {
		return
	}

	// remove me from the request and clean
	msg = strings.Replace(msg, tokmsg[0], "", 1)
	msg = strings.TrimSpace(msg)

	// pull off the first word as a command token
	cmd := strings.ToLower(tokmsg[1])
	msg = strings.Replace(msg, cmd, "", 1)
	msg = strings.TrimSpace(msg)

	// the big switch statement in the sky
	r := "Sorry, no idea what that means."
	switch cmd {
	case "help":
		r = Help(config)
	case "secrets":
		r = Secrets(config)
	case "build":
		r = Build(config.BambooUser, config.BambooPass, msg)
	case "deploy":
		r = Deploy()
	case "catchphrase":
		r = CatchPhrase()
	case "rename":
		r = Rename(config, msg)
	case "motivate":
		r = Motivate(msg)
	case "big":
		r = Big(msg)
	case "jane":
		r = Jane()
	case "hal":
		r = Hal()
	case "sensu":
		r = Sensu()
	case "rules":
		r = Rules()
	case "env":
		r = Env(config.BambooUser, config.BambooPass)
	case "environment":
		r = Env(config.BambooUser, config.BambooPass)
	case "drop":
		r = Drop(msg)
	case "feelings":
		r = Feelings()
	}

	// feedback
	Talk(config, channel, r)

}

func Help(config *configs.Config) (r string) {
	helps := []string{
		"help",
		"build [ client | cloud | admin | html ]",
		"deploy",
		"env[ironment]",
		"rules",
		"motivate <name>",
		"catchphrase",
		"big <something>",
	}
	r = "*Say things like:*\n"
	for _, help := range helps {
		r += "\t" + config.SlackerName + " " + help + "\n"
	}
	return r
}

func Secrets(config *configs.Config) (r string) {
	secrets := []string{
		"secrets",
		"jane",
		"hal",
		"rename <name>",
		"drop <something>",
		"feelings",
	}
	r = "*Quietly say things like:*\n"
	for _, secret := range secrets {
		r += "\t" + config.SlackerName + " " + secret + "\n"
	}
	return r
}

func Talk(config *configs.Config, channel string, say string) {
	api := slack.New(config.SlackToken)
	params := slack.NewPostMessageParameters()
	params.Username = config.SlackerName
	params.IconEmoji = config.SlackerEmoji
	api.PostMessage(channel, say, params)
}
