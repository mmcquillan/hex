package connectors

import (
	"log"
	"strings"
	"fmt"
	"github.com/projectjane/jane/models"
  "gopkg.in/redis.v3"
)

type Redis struct {
}

type Environment struct {
	Address string
	Password string
	DB int64
}

func (x Redis) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Redis) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if message.In.Process {
		for _, c := range connector.Commands {
			if strings.HasPrefix(strings.ToLower(message.In.Text), strings.ToLower(c.Match)) {
				environment := Environment {
					Address: connector.Server,
					Password: connector.Pass,
					DB: 0,
				}

				status := FlushDb(environment)
				log.Println(status.String())
				message.Out.Text = fmt.Sprintf("Redis Server: %s\nStatus:%s", connector.Server, status.String())
				publishMsgs <- message
				return
			}
		}
	}
}

func (x Redis) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Redis) Help(connector models.Connector) (help string) {
	help += "jane flushdb <environment> - flushes the environments redis db\n"
	return help
}

func NewClient(environment Environment) *redis.Client {
  return redis.NewClient(&redis.Options{
        Addr:     environment.Address,
        Password: environment.Password,
        DB:       environment.DB,
  })
}

func FlushDb(environment Environment) *redis.StatusCmd {
	client := NewClient(environment)
  defer client.Close()

  return client.FlushDb()
}
