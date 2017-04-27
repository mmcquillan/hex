package core

import (
	"github.com/hexbotio/hex/models"
)

func ActiveServices(config *models.Config) (cnt int) {
	cnt = 0
	for _, service := range config.Services {
		if service.Active {
			cnt += 1
		}
	}
	return cnt
}
