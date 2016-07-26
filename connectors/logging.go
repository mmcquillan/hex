package connectors

import (
	"github.com/hpcloud/tail"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
)

type Logging struct {
}

func (x Logging) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
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
				m.Routes = connector.Routes
				m.In.Source = connector.ID
				m.In.Process = false
				m.Out.Text = connector.File + ": " + c.Name
				m.Out.Detail = line.Text
				commandMsgs <- m
			}

		}
	}
}

func (x Logging) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Logging) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Logging) Help(connector models.Connector) (help string) {
	return
}
