package parse

import ()

func SimpleMatch(fValue string, mValue string) (match bool) {
	match = false
	if fValue == mValue {
		match = true
	} else if mValue == "*" {
		match = true
	}
	return match
}
