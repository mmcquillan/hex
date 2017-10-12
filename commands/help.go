package commands

import (
	"sort"
	"strings"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

type Help struct {
}

func (x Help) Act(message *models.Message, rules *map[string]models.Rule, config models.Config) {

	help := make([]string, 0)

	// pull internal help
	help = CommandHelp(config)

	// pull all help from rules
	for _, rule := range *rules {
		if rule.Active && !rule.Hide {
			if parse.Member(rule.ACL, message.Attributes["hex.user"]) || parse.Member(rule.ACL, message.Attributes["hex.channel"]) {
				if rule.Help != "" {
					help = append(help, rule.Help)
				} else {
					help = append(help, rule.Match)
				}
			}
		}
	}

	// sort, filter and de-dupe help
	helpTitle := "Help for " + config.BotName + "..."
	sort.Strings(help)
	var lasthelp = ""
	var newhelp = make([]string, 0)
	var filter = strings.TrimSpace(strings.Replace(message.Attributes["hex.input"], "help", "", 1))
	newhelp = append(newhelp, helpTitle)
	for _, h := range help {
		if strings.Contains(h, filter) || filter == "" {
			if h != lasthelp {
				newhelp = append(newhelp, " - "+h)
			}
			lasthelp = h
		}
	}

	// output help
	if len(newhelp) > 1 {
		message.Outputs = append(message.Outputs, models.Output{
			Rule:     "help",
			Response: strings.Join(newhelp[:], "\n"),
		})
	}

}
