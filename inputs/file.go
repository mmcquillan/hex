package inputs

import (
	"github.com/hpcloud/tail"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
	"os"
	"strings"
)

type File struct {
}

// Read function
func (x File) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
	if strings.ToLower(service.Config["Mode"]) == "read" || strings.ToLower(service.Config["Mode"]) == "r" {
		hostname, _ := os.Hostname()
		seek := tail.SeekInfo{Offset: 0, Whence: 2}
		t, err := tail.TailFile(service.Config["File"], tail.Config{Follow: true, Location: &seek})
		if err != nil {
			log.Print(err)
		}
		for line := range t.Lines {
			if service.Config["Filter"] != "" {
				if parse.Match(service.Config["Filter"], line.Text) {
					message := models.MakeMessage(service.Type, service.Name, hostname, service.Config["File"], line.Text)
					inputMsgs <- message
				}
			} else {
				message := models.MakeMessage(service.Type, service.Name, hostname, service.Config["File"], line.Text)
				inputMsgs <- message
			}
		}
	}
}
