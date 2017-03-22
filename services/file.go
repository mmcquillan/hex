package services

import (
	"github.com/hpcloud/tail"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
)

type File struct {
}

func (x File) Input(inputMsgs chan<- models.Message, connector models.Connector) {
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
				m.In.Tags = parse.TagAppend(connector.Tags, c.Tags)
				m.In.Process = false
				m.Out.Text = connector.File + ": " + c.Name
				m.Out.Detail = line.Text
				inputMsgs <- m
			}

		}
	}
}

func (x File) Action(message models.Message, outputMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x File) Output(outputMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x File) Help(connector models.Connector) (help []string) {
	return
}
