package connectors

import (
	"github.com/mmcquillan/jane/models"
)

type Connector interface {
	Run(config *models.Config, connector models.Connector)
	//Interpret(config *models.Config, message models.Message)
	//Send(config *models.Config, message models.Message)
}
