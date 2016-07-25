package parse

import (
	"regexp"
	"strconv"
	"strings"
)

func Match(pattern string, value string) (match bool, tokens map[string]string) {
	pattern = strings.ToLower(pattern)
	ws := strings.HasPrefix(pattern, "*")
	we := strings.HasSuffix(pattern, "*")
	re := strings.HasPrefix(pattern, "/") && strings.HasSuffix(pattern, "/")
	if re {
		pattern = strings.Replace(pattern, "/", "", -1)
	} else {
		pattern = strings.Replace(pattern, "*", "", -1)
	}
	value = strings.TrimSpace(value)
	tokens = make(map[string]string)
	tokens["0"] = value
	ta := strings.Split(strings.TrimSpace(strings.Replace(value, pattern, "", 1)), " ")
	for i := 0; i < len(ta); i++ {
		tokens[strconv.Itoa(i+1)] = ta[i]
	}
	tokens["*"] = strings.Join(ta, " ")
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
	return match, tokens
}
