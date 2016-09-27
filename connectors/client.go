package connectors

import (
	"encoding/json"
	"fmt"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
	"net"
	"time"
)

type Client struct {
}

func (x Client) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Client) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Client) Publish(publishMsgs <-chan models.Message, connector models.Connector) {

	// retry
	defer func() {
		if r := recover(); r != nil {
			log.Print("Retrying Client Connector - ", connector.ID, " ", r)
			time.Sleep(8 * time.Second)
			x.Publish(publishMsgs, connector)
		}
	}()

	// connect
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", connector.Server+":"+connector.Port)
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()

	// validate key
	key, err := parse.MakeKey(connector.Key)
	if err != nil {
		log.Print(err)
	}
	handshake, err := parse.Encrypt(key, key)
	if err != nil {
		log.Print(err)
	}
	if connector.Debug {
		log.Print("Sending Handshake")
	}
	fmt.Fprintln(conn, handshake)

	// loop
	for {

		// incoming message
		message := <-publishMsgs

		// serialize message
		messageJson, err := json.Marshal(message)
		if err != nil {
			log.Print(err)
		}

		// encrypt message
		messageEncrypt, err := parse.Encrypt(key, messageJson)
		if err != nil {
			log.Print(err)
		}

		// publish
		fmt.Fprintln(conn, messageEncrypt)

	}
	return
}

func (x Client) Help(connector models.Connector) (help []string) {
	return
}
