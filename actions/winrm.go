package actions

import (
	"bufio"
	"bytes"
	"log"
	"strconv"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
	"github.com/hexbotio/winrm"
)

//WinRM Struct representing a WinRM Connector
type WinRM struct {
}

//Action Standard command parser
func (x WinRM) Act(action models.Action, message *models.Message, config *models.Config) {

	// have to lookup each matching server
	for _, service := range config.Services {
		if service.Active && service.Type == "winrm" {
			if parse.Match(action.Service, service.Name) {

				var results = ""
				command := parse.Substitute(action.Command, message.Inputs)

				port := 5985
				if service.Config["Port"] != "" {
					port, _ = strconv.Atoi(service.Config["Port"])
				}

				endpoint := winrm.NewEndpoint(service.Config["Server"], port, false, false, nil, nil, nil, 0)
				rmclient, err := winrm.NewClient(endpoint, service.Config["Login"], service.Config["Pass"])
				if err != nil {
					log.Print("Error connecting to endpoint:", err)
					results = "Error connecting to endpoint."
				}

				var in bytes.Buffer
				var out bytes.Buffer
				var e bytes.Buffer

				stdin := bufio.NewReader(&in)
				stdout := bufio.NewWriter(&out)
				stderr := bufio.NewWriter(&e)

				_, err = rmclient.RunWithInput(command, stdout, stderr, stdin)
				if err != nil {
					log.Print("Error running command: " + err.Error())
				}

				if e.String() != "" {
					results = e.String()
				} else {
					results = out.String()
				}

				message.Response = append(message.Response, results)
				message.Success = ActionSuccess(results, action.Success, action.Failure)

			}
		}
	}
}
