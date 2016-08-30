package connectors

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/remoteExecution"
)

//WinRM Struct representing a WinRM Connector
type WinRM struct {
}

//Listen Listen not implemented for WinRM
func (x WinRM) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	remoteExecution.RemoteListen(commandMsgs, connector)
}

//Command Standard command parser
func (x WinRM) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	remoteExecution.RemoteCommand(message, publishMsgs, connector)
}

//Publish Not implemented for WinRM
func (x WinRM) Publish(connector models.Connector, message models.Message, target string) {
	return
}

//Help Returns help information for the connector
func (x WinRM) Help(connector models.Connector) (help string) {
	return remoteExecution.RemoteHelp(connector)
}
