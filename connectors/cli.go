package connectors

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/projectjane/jane/models"
	"log"
	"os"
	"os/user"
)

type Cli struct {
}

func (x Cli) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	u, err := user.Current()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Starting in cli mode...\n")
	fmt.Print("jane> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		if connector.Debug {
			log.Print("CLI Incoming message: " + req)
		}
		var m models.Message
		m.Routes = connector.Routes
		m.In.Source = connector.ID
		m.In.User = u.Username
		m.In.Text = req
		m.In.Process = true
		commandMsgs <- m
	}
}

func (x Cli) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if message.In.Text == "exit" {
		log.Print("Exiting jane by command line")
		os.Exit(0)
	}
	fmt.Println("")
	fmt.Print("\njane> ")
}

func (x Cli) Publish(connector models.Connector, message models.Message, target string) {
	switch message.Out.Status {
	case "SUCCESS":
		color.Set(color.FgGreen)
	case "WARN":
		color.Set(color.FgYellow)
	case "FAIL":
		color.Set(color.FgRed)
	}
	fmt.Println(message.Out.Text)
	color.Unset()
	if message.Out.Detail != "" {
		fmt.Println(message.Out.Detail)
	}
}

func (x Cli) Help(connector models.Connector) (help string) {
	help = "exit - close and exit\n"
	return help
}
