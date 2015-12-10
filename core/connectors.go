package core

import (
	"github.com/projectjane/jane/models"
)

func ActiveConnectors(config *models.Config) (cnt int) {
	cnt = 0
	for _, connector := range config.Connectors {
		if connector.Active {
			cnt += 1
		}
	}
	return cnt
}
