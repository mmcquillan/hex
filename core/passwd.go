package core

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/projectjane/jane/models"
)

//Passwd Generates a random password from aes-256 key
func Passwd(message models.Message, publishMsgs chan<- models.Message) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		log.Printf("Passwd Error: %s", err)
		message.Out.Text = "Error generating password"
	} else {
		pass := base64.StdEncoding.EncodeToString([]byte(key))
		message.Out.Text = pass
	}

	publishMsgs <- message
}
