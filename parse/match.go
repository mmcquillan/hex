package parse

import (
	"regexp"
	"strings"
)

func Match(pattern string, value string) (match bool) {
	ws := strings.HasPrefix(pattern, "*")
	we := strings.HasSuffix(pattern, "*")
	re := strings.HasPrefix(pattern, "/") && strings.HasSuffix(pattern, "/")
	if re {
		pattern = strings.Replace(pattern, "/", "", -1)
	} else {
		pattern = strings.Replace(pattern, "*", "", -1)
	}
	value = strings.TrimSpace(value)
	if !re {
		pattern = strings.ToLower(pattern)
	}
	value = strings.ToLower(value)
	match = false
	if re {
		regx := regexp.MustCompile(pattern)
		match = regx.MatchString(value)
	} else if ws && we && strings.Contains(value, pattern) {
		match = true
	} else if ws && !we && strings.HasSuffix(value, pattern) {
		match = true
	} else if !ws && we && strings.HasPrefix(value, pattern) {
		match = true
	} else if value == pattern {
		match = true
	}
	return match
}
