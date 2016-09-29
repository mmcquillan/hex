package connectors

import (
	"bufio"
	"encoding/json"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type Server struct {
}

func (x Server) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)

	// estsablish connection
	clients := make(map[string]int)
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

	// watch for stopped connections
	go func(commandMsgs chan<- models.Message, connector models.Connector, clients map[string]int) {
		for {
			time.Sleep(15 * time.Second)
			for client, _ := range clients {
				clients[client] -= 1
				if clients[client] < 0 {
					var m models.Message
					m.In.ConnectorType = connector.Type
					m.In.ConnectorID = connector.ID
					m.In.Tags = connector.Tags
					m.In.Process = false
					m.Out.Text = "Client Alert"
					m.Out.Detail = "Disconnect from " + client
					m.Out.Status = "FAIL"
					commandMsgs <- m
					if connector.Debug {
						log.Print("Disconnect: ", client)
					}
					delete(clients, client)
				}
			}
		}
	}(commandMsgs, connector, clients)

	for {

		// respond to connections
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		if connector.Debug {
			log.Print("New Client Connection by ", conn.RemoteAddr())
		}

		// thread out process
		go func(commandMsgs chan<- models.Message, connector models.Connector, conn net.Conn, key []byte, clients map[string]int) {
			defer conn.Close()

			// establish client counter
			client := strings.Split(conn.RemoteAddr().String(), ":")[0]
			clients[client] = 3
			var m models.Message
			m.In.ConnectorType = connector.Type
			m.In.ConnectorID = connector.ID
			m.In.Tags = connector.Tags
			m.In.Process = false
			m.Out.Text = "Client Alert"
			m.Out.Detail = "Connect from " + client
			m.Out.Status = "SUCCESS"
			commandMsgs <- m

			// loop
			for {

				// read in messages
				messageEvent, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					if err == io.EOF {
						clients[client] = 0
						break
					} else {
						log.Print(err)
					}
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

				// interpret message
				if string(messageDecrypt) == "ping" {
					clients[client] += 1
				} else if len(messageDecrypt) > 0 && string(messageDecrypt)[0:1] == "{" {

					// deserialize message
					var message models.Message
					err = json.Unmarshal(messageDecrypt, &message)
					if err != nil {
						log.Print(err)
					}
					message.In.ConnectorID = "[" + client + "]" + message.In.ConnectorID
					message.In.Tags = parse.TagAppend(message.In.Tags, connector.Tags)
					if connector.Debug {
						log.Printf("Message from %s: %+v", client, message)
					}

					// send message
					commandMsgs <- message

				} else {
					log.Print("Client Connection illegal key from ", client)
					clients[client] = 0
					break
				}

			}
		}(commandMsgs, connector, conn, key, clients)
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
