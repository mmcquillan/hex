package commands

import (
	"bytes"
	"log"
	"os/exec"
)

func Big(msg string) (results string) {
	cmd := exec.Command("/usr/bin/figlet", "-w80", msg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	results = "```" + out.String() + "```"
	return results
}
