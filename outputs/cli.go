package outputs

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/hexbotio/hex/models"
)

type Cli struct {
}

func (x Cli) Write(message models.Message, config models.Config) {
	fmt.Print("\n")
	for _, output := range message.Outputs {
		if output.State == "" {
			fmt.Println(output.Response, "\n")
		} else {
			color.Set(color.FgBlue)
			switch output.State {
			case models.PASS:
				color.Set(color.FgGreen)
			case models.WARN:
				color.Set(color.FgYellow)
			case models.FAIL:
				color.Set(color.FgRed)
			}
			fmt.Println(output.Response, "\n")
			color.Unset()
		}
	}
	if message.Debug {
		keys := make([]string, 0, len(message.Attributes))
		for key, _ := range message.Attributes {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		fmt.Printf("MESSAGE DEBUG (%d sec to complete)\n", models.MessageTimestamp()-message.CreateTime)
		for _, key := range keys {
			fmt.Printf("  %s: '%s'\n", key, message.Attributes[key])
		}
	}
	fmt.Print("\n", config.BotName, "> ")
}
