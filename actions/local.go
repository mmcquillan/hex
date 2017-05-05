package actions

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/hexbotio/hex/models"
	"github.com/hexbotio/hex/parse"
)

//Exec struct
type Local struct {
}

//Action function
func (x Local) Act(action models.Action, message *models.Message, config *models.Config) {
	var o bytes.Buffer
	var e bytes.Buffer
	command := parse.Substitute(action.Command, message.Inputs)
	c := exec.Command("/bin/sh", "-c", command)
	path := message.Inputs["hex.pipeline.workspace"]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	c.Dir = path
	c.Stdout = &o
	c.Stderr = &e
	err := c.Run()
	if err != nil {
		log.Print(action.Command)
		log.Print(err)
		log.Print(e.String())
	}
	out := o.String()
	message.Response = append(message.Response, out)
	message.Success = ActionSuccess(out, action.Success, action.Failure)
}
