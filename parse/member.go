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

func EitherMember(groups string, member1 string, member2 string) (match bool) {
	if Member(groups, member1) || Member(groups, member2) {
		return true
	}
	return false
}
