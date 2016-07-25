package parse

import (
	"os"
	"regexp"
	"strings"
)

func Substitute(value string, tokens map[string]string) (result string) {
	if match, hits := SubstitutionVars(value); match {
		for _, hit := range hits {
			if _, ok := tokens[Strip(hit)]; ok {
				value = strings.Replace(value, hit, tokens[Strip(hit)], -1)
			} else {
				value = strings.Replace(value, hit, os.Getenv(Strip(hit)), -1)
			}
		}
	}
	return value
}

func SubstitutionVars(value string) (match bool, tokens []string) {
	match = false
	re := regexp.MustCompile("\\${([A-Za-z0-9*_\\-\\.]+)}")
	tokens = re.FindAllString(value, -1)
	if len(tokens) > 0 {
		match = true
	}
	return match, tokens
}

func Strip(value string) (stripped string) {
	stripped = strings.Replace(value, "${", "", -1)
	stripped = strings.Replace(stripped, "}", "", -1)
	return stripped
}
