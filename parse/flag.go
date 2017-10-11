package parse

import (
	"strings"
)

func Flag(input string, flag string) (cleaned_input string, flag_match bool) {
	cleaned_input = strings.TrimSpace(input)
	if strings.HasSuffix(cleaned_input, " "+flag) {
		return strings.TrimSpace(strings.TrimRight(cleaned_input, flag)), true
	}
	return input, false
}
