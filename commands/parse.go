package commands

import (
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/outputs"
	"strings"
)

func Parse(config *configs.Config, channel string, msg string) {

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
		r = Build(config.BambooUrl, config.BambooUser, config.BambooPass, msg)
	case "deploy":
		r = Deploy(config.BambooUrl)
	case "rename":
		r = Rename(config, msg)
	case "big":
		r = Big(msg)
	case "sensu":
		r = Sensu()
	case "env":
		r = Env(config.BambooUrl, config.BambooUser, config.BambooPass)
	case "environment":
		r = Env(config.BambooUrl, config.BambooUser, config.BambooPass)
	case "reload":
		r = Reload(config)
	default:
		r = Response(config, cmd, msg)
	}

	// feedback
	message := outputs.Message{channel, r, "", "", ""}
	outputs.Output(config, message)

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
		"rename <name>",
		"reload",
	}
	r = "*Quietly say things like:*\n"
	for _, secret := range secrets {
		r += "\t" + config.Name + " " + secret + "\n"
	}
	return r
}
