package connectors

import (
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/remoteExecution"
)

//Exec Struct representing an exec connector
type Exec struct {
}

//Listen Exec listen implementation
func (x Exec) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	remoteExecution.RemoteListen(commandMsgs, connector)
}

//Command Exec command implementation
func (x Exec) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	remoteExecution.RemoteCommand(message, publishMsgs, connector)
}

//Publish Exec publish implementation
func (x Exec) Publish(connector models.Connector, message models.Message, target string) {
	return
}

//Help Exec help implementation
func (x Exec) Help(connector models.Connector) (help string) {
	return remoteExecution.RemoteHelp(connector)
}
