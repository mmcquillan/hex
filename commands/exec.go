package commands

import (
	"bytes"
	"github.com/mmcquillan/jane/models"
	"log"
	"os/exec"
	"strings"
)

func Exec(msg string, command models.Command) (results string) {
	msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
	msg = strings.Replace(msg, "\"", "", -1)
	args := strings.Split(strings.Replace(command.Args, "%msg%", msg, -1), " ")
	cmd := exec.Command(command.Cmd, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Print(command.Cmd + " " + strings.Join(args, " "))
		log.Print(err)
	}
	results = strings.Replace(command.Output, "%stdout%", out.String(), -1)
	return results
}
