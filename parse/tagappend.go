package parse

import (
	"strings"
)

func TagAppend(oldTags string, newTags string) (allTags string) {
	if oldTags != "" && newTags != "" {
		tagList := make([]string, 0)
		for _, t := range strings.Split(oldTags+","+newTags, ",") {
			if t != "" {
				tagList = append(tagList, t)
			}
		}
		allTags = strings.Join(tagList, ",")
	} else if oldTags != "" && newTags == "" {
		allTags = oldTags
	} else if oldTags == "" && newTags != "" {
		allTags = newTags
	} else {
		allTags = ""
	}
	return allTags
}
