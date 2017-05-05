package actions

import (
	"log"
	"strconv"
	"time"

	"github.com/hexbotio/hex/models"
)

//Exec struct
type Sleep struct {
}

//Action function
func (x Sleep) Act(action models.Action, message *models.Message, config *models.Config) {
	wait, err := strconv.Atoi(action.Command)
	if err != nil {
		log.Print("ERROR - Cannot parse Command as Int in action:" + action.Name)
	}
	time.Sleep(time.Duration(wait) * time.Millisecond)
}
