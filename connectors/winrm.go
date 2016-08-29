package connectors

import (
	"bufio"
	"bytes"
	"errors"
	"log"

	"github.com/masterzen/winrm"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
)

//WinRM Struct representing a WinRM Connector
type WinRM struct {
	Client *winrm.Client
}

//Listen Listen not implemented for WinRM
func (x WinRM) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
}

//Command Standard command parser
func (x WinRM) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if match, tokens := parse.Match("winrm*", message.In.Text); match {
		command := tokens["*"]

		if x.Client == nil {
			log.Println("Client is nil. Connecting...")

			var err error

			endpoint := winrm.NewEndpoint(connector.Server, 5985, false, false, nil, nil, nil, 0)
			x.Client, err = winrm.NewClient(endpoint, connector.Login, connector.Pass)
			if err != nil {
				log.Println("Error connecting to endpoint:", err)
			}
		}

		out, err := x.sendCommand(command)
		if err != nil {
			log.Println("Error sending command:", err)
			message.Out.Text = "Error processing command: " + err.Error()
			message.Out.Status = "FAIL"
		} else {
			message.Out.Text = out
			message.Out.Status = "SUCCESS"
		}

		publishMsgs <- message
	}
}

//Publish Not implemented for WinRM
func (x WinRM) Publish(connector models.Connector, message models.Message, target string) {
	return
}

//Help Returns help information for the connector
func (x WinRM) Help(connector models.Connector) (help string) {
	help += "winrm {cmd}"
	return help
}

func (x WinRM) sendCommand(command string) (string, error) {
	var in bytes.Buffer
	var out bytes.Buffer
	var e bytes.Buffer

	stdin := bufio.NewReader(&in)
	stdout := bufio.NewWriter(&out)
	stderr := bufio.NewWriter(&e)

	_, err := x.Client.RunWithInput(command, stdout, stderr, stdin)
	if err != nil {
		return "", err
	}

	if e.String() != "" {
		return "", errors.New(e.String())
	}

	return out.String(), nil
}
