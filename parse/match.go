package parse

import (
	"regexp"
	"strconv"
	"strings"
)

func Match(pattern string, value string) (match bool, tokens map[string]string) {
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
	tokens["0"] = pattern
	if !re {
		pattern = strings.ToLower(pattern)
	}
	ta := strings.Split(CIReplace(value, pattern, ""), " ")
	for i := 0; i < len(ta); i++ {
		tokens[strconv.Itoa(i+1)] = strings.TrimSpace(ta[i])
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

func CIReplace(str string, val string, rep string) (ret string) {
	f := strings.Index(strings.ToLower(str), strings.ToLower(val))
	if f == -1 {
		ret = str
	} else {
		l := len(val)
		ret = strings.TrimSpace(str[:f]) + " " + strings.TrimSpace(str[(f+l):])
	}
	return strings.TrimSpace(ret)
}
