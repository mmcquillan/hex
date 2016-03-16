package connectors

import (
	"log"
	"strings"
	"encoding/json"
	"github.com/projectjane/jane/models"
  "gopkg.in/redis.v3"
)

type Redis struct {
	Environments map[string][]Environment
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
	if strings.HasPrefix(strings.ToLower(message.In.Text), "flushdb") {
		env := strings.TrimSpace(strings.Replace(message.In.Text, "flusbdb", "", 1))
		environment := LookupEnvironment(env)
		status := FlushDb(environment)
		log.Println(status.String())
		message.Out.Text = status.String()
		publishMsgs <- message
	}
}

func (x Redis) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Redis) Help(connector models.Connector) (help string) {
	help += "flushdb <environment> - pulls back an image url\n"
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

func (redis Redis) LookupEnvironment(env string) Environment {
	return redis.Environments[env]
}

func (x Redis) GetEnvironmentsFromConfig(config models.Config) {
	var envs []Environment
	err := json.Unmarshal([]byte(config), &envs)
	if err != nil {
		log.Println(err)
	}
}
