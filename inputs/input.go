package inputs

import (
	"bufio"
	"fmt"
	"github.com/mmcquillan/jane/commands"
	"github.com/mmcquillan/jane/configs"
	"github.com/nlopes/slack"
	"os"
)

func Input(config *configs.Config) {
	if config.Interactive {
		fmt.Println("Jane starting in interactive mode...\n")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println("")
			commands.Parse(config, "", scanner.Text())
		}
	} else {
		api := slack.New(config.SlackToken)
		api.SetDebug(false)
		rtm := api.NewRTM()
		go rtm.ManageConnection()
		for {
			select {
			case msg := <-rtm.IncomingEvents:
				switch ev := msg.Data.(type) {
				case *slack.MessageEvent:
					if ev.User != "" {
						commands.Parse(config, ev.Channel, ev.Text)
					}
				}
			}
		}
	}

}
