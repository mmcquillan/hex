package connectors

import (
	"github.com/projectjane/jane/models"
	"log"
  "gopkg.in/redis.v3"
)

type Redis struct {
}

func (x Template) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Template) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if strings.Contains(strings.ToLower(message.In.Text), "flushdb") {
		status := FlushDb()
		message.Out.Text = string(status
		publishMsgs <- message)
	}

	return
}

func (x Template) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Template) Help(connector models.Connector) (help string) {
	return
}

func NewClient(addr, pass string) *redis.Client {
  return redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: pass,
        DB:       0,  // use default DB
  })
}

func FlushDb(addr, pass string) *redis.StatusCmd {
	client := NewClient(addr, pass)
  defer client.Close()

  return client.FlushDb()
}
