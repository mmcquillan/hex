package relays

import (
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/configs"
	"strings"
)

func Parse(config *configs.Config, relay configs.Relay, channel string, msg string) {

	// make sure they are talking to and not about us
	tokmsg := strings.Split(strings.TrimSpace(msg), " ")
	if strings.ToLower(tokmsg[0]) != strings.ToLower(config.Name) {
		return
	}

	// remove me from the request and clean
	msg = strings.Replace(msg, tokmsg[0], "", 1)
	msg = strings.TrimSpace(msg)

	// see if nothing is said
	if msg == "" {
		return
	}

	// pull off the first word as a command token
	cmd := strings.ToLower(tokmsg[1])
	msg = strings.Replace(msg, cmd, "", 1)
	msg = strings.TrimSpace(msg)

	// the big switch statement in the sky
	var r string
	switch cmd {
	case "help":
		r = Help(config)
	case "secrets":
		r = Secrets(config)
	case "build":
		r = commands.Build(config.BambooUrl, config.BambooUser, config.BambooPass, msg)
	case "deploy":
		r = commands.Deploy(config.BambooUrl)
	case "big":
		r = commands.Big(msg)
	case "sensu":
		r = commands.Sensu()
	case "env":
		r = commands.Env(config.BambooUrl, config.BambooUser, config.BambooPass)
	case "environment":
		r = commands.Env(config.BambooUrl, config.BambooUser, config.BambooPass)
	case "reload":
		r = commands.Reload(config)
	default:
		r = commands.Response(config, cmd, msg)
	}

	// feedback
	message := Message{channel, r, "", "", ""}
	Output(config, relay, message)

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
		r += "\t" + config.Name + " " + help + "\n"
	}
	return r
}

func Secrets(config *configs.Config) (r string) {
	secrets := []string{
		"secrets",
		"reload",
	}
	r = "*Quietly say things like:*\n"
	for _, secret := range secrets {
		r += "\t" + config.Name + " " + secret + "\n"
	}
	return r
}
