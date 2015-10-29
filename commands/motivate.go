package commands

import (
	"strings"
)

func Motivate(msg string) (results string) {
	name := strings.TrimSpace(msg)
	results = "You can _do it_ " + name + "!"
	return results
}
