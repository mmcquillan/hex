package core

import (
	"os"
	"strings"
)

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func DirExists(dir string) bool {
	return FileExists(dir)
}

func MakePath(dir string, file string) string {
	if strings.HasSuffix(dir, "/") {
		return dir + file
	}
	return dir + "/" + file
}
