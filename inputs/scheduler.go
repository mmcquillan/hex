package inputs

import (
	"sync"

	"github.com/hexbotio/hex/models"
	"github.com/robfig/cron"
)

//Exec struct
type Scheduler struct {
}

//Input function
func (x Scheduler) Read(inputMsgs chan<- models.Message, config models.Config) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	cron := cron.New()
	/*
		cron.AddFunc(service.Config["Schedule"], func() {
			message := models.MakeMessage(service.Type, service.Name, service.Config["Schedule"], "", "")
			inputMsgs <- message
		})
	*/
	cron.Start()
	defer cron.Stop()
	wg.Wait()

	config.Logger.Warn("Scheduler Ending")
}
