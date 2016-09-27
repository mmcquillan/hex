package connectors

import (
	"bufio"
	"encoding/json"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"log"
	"net"
)

type Server struct {
}

func (x Server) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	// estsablish connection
	listener, err := net.Listen("tcp", ":"+connector.Port)
	if err != nil {
		log.Print(err)

	}
	defer listener.Close()

	// validate key
	key, err := parse.MakeKey(connector.Key)
	if err != nil {
		log.Print(err)
	}

	for {

		// respond to connections
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		if connector.Debug {
			log.Print("New Client Connection by ", conn.RemoteAddr())
		}

		go func(commandMsgs chan<- models.Message, connector models.Connector, conn net.Conn, key []byte) {
			defer conn.Close()

			// first message should be key to validate
			handshake, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Print(err)
			}
			handshakeBytes, err := parse.StrBytesToBytes(handshake)
			if err != nil {
				log.Print(err)
			}
			handshakeDecrypt, err := parse.Decrypt(key, handshakeBytes)
			if err != nil {
				log.Print(err)

			}
			if string(handshakeDecrypt) != string(key) {
				log.Printf("Handshake: %+v", handshakeDecrypt)
				log.Printf("Key: %+v", key)
				log.Print("Key is invalid from client")
				return
			} else {
				log.Print("Handshake success")
			}

			// loop
			for {

				// read in messages
				messageEvent, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Print(err)
				}

				// convert to bytes
				messageBytes, err := parse.StrBytesToBytes(messageEvent)
				if err != nil {
					log.Print(err)
				}

				// decrypt message
				messageDecrypt, err := parse.Decrypt(key, messageBytes)
				if err != nil {
					log.Print(err)
				}

				// deserialize message
				var message models.Message
				err = json.Unmarshal(messageDecrypt, &message)
				if err != nil {
					log.Print(err)
				}
				message.In.ConnectorID = "[" + conn.RemoteAddr().String() + "]" + message.In.ConnectorID
				message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags)
				if connector.Debug {
					log.Printf("Message from %s: %+v", conn.RemoteAddr().String(), message)
				}

				// send message
				//commandMsgs <- message

			}
		}(commandMsgs, connector, conn, key)
	}

	return
}

func (x Server) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	return
}

func (x Server) Publish(publishMsgs <-chan models.Message, connector models.Connector) {
	return
}

func (x Server) Help(connector models.Connector) (help []string) {
	return
}
