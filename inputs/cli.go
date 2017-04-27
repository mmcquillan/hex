package inputs

import (
	"bufio"
	"fmt"
	"github.com/projectjane/jane/models"
	"os"
	"os/user"
	"strings"
)

// Cli struct
type Cli struct {
}

// Read function
func (x Cli) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
	hostname, _ := os.Hostname()
	user, _ := user.Current()
	fmt.Println("Starting in cli mode...\n")
	fmt.Print(service.BotName, "> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		req := scanner.Text()
		fmt.Print("\n", service.BotName, "> ")
		if strings.TrimSpace(req) != "" {
			message := models.MakeMessage(service.Type, service.Name, hostname, user.Username, req)
			inputMsgs <- message
		}
	}
}
