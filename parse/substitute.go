package parse

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
)

//Substitute Function for substituting a string with tokens
func Substitute(value string, tokens map[string]string) string {
	if match, hits := findVars(value); match {
		for _, hit := range hits {
			clean_hit := strip(hit)
			if strings.HasPrefix(clean_hit, "hex.input.json:") {
				value = strings.Replace(value, hit, subJson(clean_hit, tokens["hex.input"]), -1)
			} else if strings.HasPrefix(clean_hit, "hex.input.") {
				value = strings.Replace(value, hit, subInput(clean_hit, tokens), -1)
			} else if _, ok := tokens[clean_hit]; ok {
				value = strings.Replace(value, hit, tokens[clean_hit], -1)
			} else {
				value = strings.Replace(value, hit, os.Getenv(clean_hit), -1)
			}
		}
	}
	return value
}

func SubstituteEnv(value string) string {
	if match, hits := findVars(value); match {
		for _, hit := range hits {
			value = strings.Replace(value, hit, os.Getenv(strip(hit)), -1)
		}
	}
	return value
}

func findVars(value string) (match bool, tokens []string) {
	match = false
	re := regexp.MustCompile("\\${([A-Za-z0-9:*_\\-\\.\\?]+)}")
	tokens = re.FindAllString(strings.Replace(value, "$${", "X{", -1), -1)
	if len(tokens) > 0 {
		match = true
	}
	return match, tokens
}

func subJson(token string, json string) (out string) {
	jsonParsed, err := gabs.ParseJSON([]byte(json))
	if err != nil {
		return out
	}
	token = strings.Replace(token, "hex.input.json:", "", -1)
	value, ok := jsonParsed.Path(token).Data().(string)
	if ok {
		out = value
	}
	return out
}

func subInput(input string, tokens map[string]string) (out string) {
	tokenInput := strings.Split(tokens["hex.input"], " ")
	inputEval := strings.Replace(input, "hex.input.", "", -1)
	var tokenStart int
	var tokenEnd int
	var err error
	inputRange := strings.Split(inputEval, ":")
	if len(inputRange) == 1 {
		if inputRange[0] == "*" {
			out = tokens["hex.input"]
		}
		tokenStart, err = strconv.Atoi(inputRange[0])
		if err != nil {
			return out
		}
		if tokenStart >= len(tokenInput) {
			return out
		}
		tokenEnd = tokenStart + 1
	}
	if len(inputRange) == 2 {
		tokenStart, err = strconv.Atoi(inputRange[0])
		if err != nil {
			return out
		}
		if inputRange[1] == "*" {
			tokenEnd = len(tokenInput)
		} else {
			tokenEnd, err = strconv.Atoi(inputRange[1])
			if err != nil {
				return out
			}
			tokenEnd = tokenEnd + 1
			if tokenEnd > len(tokenInput) {
				tokenEnd = len(tokenInput)
			}
		}
		if tokenStart > tokenEnd {
			return out
		}
	}
	out = strings.Join(tokenInput[tokenStart:tokenEnd], " ")
	return out
}

func strip(value string) (stripped string) {
	stripped = strings.Replace(value, "${", "", -1)
	stripped = strings.Replace(stripped, "}", "", -1)
	return stripped
}
