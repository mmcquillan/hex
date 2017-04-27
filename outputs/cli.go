package outputs

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/projectjane/jane/models"
)

type Cli struct {
}

func (x Cli) Write(outputMsgs <-chan models.Message, service models.Service) {
	for {
		message := <-outputMsgs
		fmt.Print("\n")
		if message.Success {
			color.Set(color.FgGreen)
		} else {
			color.Set(color.FgRed)
		}
		fmt.Println(strings.Join(message.Response[:], "\n"))
		color.Unset()
		fmt.Print("\n", service.BotName, "> ")
	}
}
