package inputs

import (
	"log"
	"sync"

	"github.com/hexbotio/hex/models"
	"github.com/robfig/cron"
)

//Exec struct
type Scheduler struct {
}

//Input function
func (x Scheduler) Read(inputMsgs chan<- models.Message, service models.Service) {
	defer Recovery(service)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	cron := cron.New()
	cron.AddFunc(service.Config["Schedule"], func() {
		message := models.MakeMessage(service.Type, service.Name, service.Config["Schedule"], "", "")
		inputMsgs <- message
	})
	cron.Start()
	defer cron.Stop()
	wg.Wait()

	log.Print("ending")
}
