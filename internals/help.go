package internals

import (
	"github.com/hexbotio/hex/models"
	"sort"
	"strings"
)

type Help struct {
}

func (x Help) Act(message *models.Message, config *models.Config) {

	help := make([]string, 0)

	// pull internal help
	help = InternalHelp(config)

	// pull all help from the aliases
	for _, alias := range config.Aliases {
		if !alias.Hide {
			if alias.Help != "" {
				help = append(help, alias.Help)
			} else {
				help = append(help, alias.Match)
			}
		}
	}

	// pull all help from the connectors
	for _, pipeline := range config.Pipelines {
		if pipeline.Active {
			for _, input := range pipeline.Inputs {
				if input.Type == message.Inputs["hex.type"] || input.Type == "*" {
					if !(input.Hide || input.Match == "*") {
						if input.Help != "" {
							help = append(help, input.Help)
						} else {
							help = append(help, input.Match)
						}
					}
				}
			}
		}
	}

	// sort, filter and de-dupe help
	helpTitle := "Help for " + config.BotName + "..."
	sort.Strings(help)
	var lasthelp = ""
	var newhelp = make([]string, 0)
	var filter = strings.TrimSpace(strings.Replace(message.Inputs["hex.input"], config.BotName+" help", "", 1))
	newhelp = append(newhelp, helpTitle)
	for _, h := range help {
		if strings.Contains(h, filter) || filter == "" {
			if h != lasthelp {
				newhelp = append(newhelp, "* "+h)
			}
			lasthelp = h
		}
	}

	// output help
	if len(newhelp) > 1 {
		message.Response = newhelp
	}

}
