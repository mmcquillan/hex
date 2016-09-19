package connectors

import (
	"fmt"
	"github.com/projectjane/jane/models"
	"github.com/projectjane/jane/parse"
	"gopkg.in/redis.v3"
	"log"
)

type Redis struct {
}

type Environment struct {
	Address  string
	Password string
	DB       int64
}

func (x Redis) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	return
}

func (x Redis) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	for _, c := range connector.Commands {
		if match, _ := parse.Match(c.Match, message.In.Text); match {
			environment := Environment{
				Address:  connector.Server,
				Password: connector.Pass,
				DB:       0,
			}

			status := FlushDb(environment)
			log.Println(status.String())
			message.In.Tags += "," + connector.Tags
			message.Out.Text = fmt.Sprintf("Redis Server: %s\nStatus:%s", connector.Server, status.String())
			publishMsgs <- message
			return
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
