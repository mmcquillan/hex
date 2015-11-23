package connectors

import (
	"github.com/projectjane/jane/models"
)

type Connector interface {
	Listen(config *models.Config, connector models.Connector)
	Command(config *models.Config, message *models.Message)
	Publish(config *models.Config, connector models.Connector, message models.Message, target string)
}
