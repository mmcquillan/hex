package commands

import (
	"github.com/projectjane/jane/models"
	"strings"
)

func Response(msg string, command models.Command) (results string) {
	msg = strings.TrimSpace(strings.Replace(msg, command.Match, "", 1))
	results = strings.Replace(command.Output, "%msg%", msg, -1)
	return results
}
