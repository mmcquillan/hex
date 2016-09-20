package parse

import (
	"strings"
)

func TagMatch(findTags string, matchTags string) (match bool) {
	match = false
	if matchTags == "*" {
		match = true
	} else {
		fTags := strings.Split(findTags, ",")
		mTags := strings.Split(matchTags, ",")
		for _, f := range fTags {
			for _, m := range mTags {
				if f == m {
					match = true
				}
			}
		}
	}
	return match
}
