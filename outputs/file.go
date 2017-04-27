package outputs

import (
	"github.com/projectjane/jane/models"
	"log"
	"os"
	"strings"
	"time"
)

type File struct {
}

func (x File) Write(outputMsgs <-chan models.Message, service models.Service) {
	if strings.ToLower(service.Config["Mode"]) == "write" || strings.ToLower(service.Config["Mode"]) == "w" {
		file, err := os.OpenFile(service.Config["File"], os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Print(err)
		}
		defer file.Close()
		for {
			message := <-outputMsgs
			if _, err = file.WriteString(time.Now().Format(time.RFC3339) + " " + strings.Join(message.Response[:], "\n") + "\n"); err != nil {
				log.Print(err)
			}
		}
	}
}
