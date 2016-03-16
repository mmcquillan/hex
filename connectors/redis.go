package connectors

import (
	"log"
	"strings"
	"github.com/projectjane/jane/models"
  "gopkg.in/redis.v3"
)

type Redis struct {
}

func (x Redis) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Redis) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if strings.Contains(strings.ToLower(message.In.Text), "flushdb") {
		status := FlushDb("", "", 0)
		log.Println(status.String())
		message.Out.Text = status.String()
		publishMsgs <- message
	}
}

func (x Redis) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Redis) Help(connector models.Connector) (help string) {
	return
}

func NewClient(addr, pass string, db int64) *redis.Client {
  return redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: pass,
        DB:       db,  // use default DB
  })
}

func FlushDb(addr, pass string, db int64) *redis.StatusCmd {
	client := NewClient(addr, pass, db)
  defer client.Close()

  return client.FlushDb()
}
