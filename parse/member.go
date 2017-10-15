package parse

import (
	"strings"
)

func Member(groups string, member string) (match bool) {
	match = false
	for _, group := range strings.Split(groups, ",") {
		if group == member || group == "*" {
			match = true
		}
	}
	return match
}
