package internals

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/projectjane/jane/models"
)

type Passwd struct {
}

func (x Passwd) Act(message *models.Message, config *models.Config) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Printf("Passwd Error: %s", err)
		message.Response = append(message.Response, "Error generating password!")
		message.Success = false
	} else {
		pass := base64.StdEncoding.EncodeToString([]byte(key))
		message.Response = append(message.Response, pass)
	}
}
