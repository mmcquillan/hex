package parse

import (
	"regexp"
	"strings"
)

func Substitution(value string) (match bool, tokens []string) {
	match = false
	re := regexp.MustCompile("{([^\\s]*)}")
	tokens = re.FindAllString(value, -1)
	if len(tokens) > 0 {
		match = true
	}
	return match, tokens
}

func Strip(value string) (stripped string) {
	stripped = strings.Replace(value, "{", "", -1)
	stripped = strings.Replace(stripped, "}", "", -1)
	return stripped
}
