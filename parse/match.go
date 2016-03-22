package parse

import (
	"strings"
)

func Match(pattern string, value string) (match bool, tokens []string) {
	pattern = strings.ToLower(pattern)
	ws := strings.HasPrefix(pattern, "*")
	we := strings.HasSuffix(pattern, "*")
	pattern = strings.Replace(pattern, "*", "", -1)
	value = strings.TrimSpace(value)
	tokens = strings.Split(strings.TrimSpace(strings.Replace(value, pattern, "", 1)), " ")
	value = strings.ToLower(value)
	match = false
	if ws && ws && strings.Contains(value, pattern) {
		match = true
	} else if ws && !we && strings.HasSuffix(value, pattern) {
		match = true
	} else if !ws && we && strings.HasPrefix(value, pattern) {
		match = true
	} else if value == pattern {
		match = true
	}
	return match, tokens
}
