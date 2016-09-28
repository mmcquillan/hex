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
	Conn      net.Conn
	Connected bool
	Connector models.Connector
}

var client Client

func (x Client) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Client) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Client) Publish(publishMsgs <-chan models.Message, connector models.Connector) {

	// init
	client = x
	client.Connector = connector
	client.Connected = false
	//defer client.Conn.Close()

	// listen for messages
	go func(publishMsgs <-chan models.Message) {

		for {

			// incoming message
			message := <-publishMsgs

			// serialize message
			messageJson, err := json.Marshal(message)
			if err != nil {
				log.Print(err)
			}

			// publish
			sendMessage(messageJson)

		}

	}(publishMsgs)

	// ping
	for {

		// send ping
		sendMessage([]byte("ping"))

		// wait
		time.Sleep(15 * time.Second)

	}

	return
}

func (x Client) Help(connector models.Connector) (help []string) {
	return
}

func sendMessage(messageBytes []byte) {

	// retry
	defer func() {
		if r := recover(); r != nil {
			client.Connected = false
			log.Print("Retrying Client Connector - ", client.Connector.ID, " ", r)
			time.Sleep(8 * time.Second)
			sendMessage(messageBytes)
		}
	}()

	// connect
	log.Print("Connected: ", client.Connected)
	if !client.Connected {
		if client.Connector.Debug {
			log.Print("Connecting to: " + client.Connector.Server + ":" + client.Connector.Port)
		}
		var err error
		client.Conn, err = net.Dial("tcp", client.Connector.Server+":"+client.Connector.Port)
		if err != nil {
			log.Print(err)
		}
		client.Connected = true
	}

	// validate key
	key, err := parse.MakeKey(client.Connector.Key)
	if err != nil {
		log.Print(err)
	}

	// encrypt message
	messageEncrypt, err := parse.Encrypt(key, messageBytes)
	if err != nil {
		log.Print(err)
	}

	// publish
	_, err = fmt.Fprintln(client.Conn, messageEncrypt)
	if err != nil {
		log.Print(err)
		client.Connected = false
		panic(err)
	}

}
