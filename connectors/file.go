package connectors

import (
	"github.com/hpcloud/tail"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
)

type File struct {
}

func (x File) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	seek := tail.SeekInfo{Offset: 0, Whence: 2}
	t, err := tail.TailFile(connector.File, tail.Config{Follow: true, Location: &seek})
	if err != nil {
		log.Print(err)
	}
	for line := range t.Lines {
		for _, c := range connector.Commands {
			if match, _ := parse.Match(c.Match, line.Text); match {
				var m models.Message
				m.In.ConnectorType = connector.Type
				m.In.ConnectorID = connector.ID
				m.In.Tags = connector.Tags
				m.In.Process = false
				m.Out.Text = connector.File + ": " + c.Name
				m.Out.Detail = line.Text
				commandMsgs <- m
			}

		}
	}
}

func (x File) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x File) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x File) Help(connector models.Connector) (help string) {
	return
}
