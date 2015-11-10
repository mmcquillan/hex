package connectors

import (
	"github.com/mmcquillan/jane/models"
)

type Connector interface {
	Run(config *models.Config, connector models.Connector)
	Send(config *models.Config, message models.Message, target string)
}
