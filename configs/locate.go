package configs

import (
	"github.com/kardianos/osext"
	"github.com/mitchellh/go-homedir"
	"os"
)

func Locate() (configFile string) {

	// order of finding the config file
	// 1. running path "./jane.config"
	// 2. users home path "~/jane.config"
	// 3. system etc "/etc/jane.config
	file := "jane.config"

	first, _ := osext.ExecutableFolder()
	first += "/" + file
	if FileExists(first) {
		return first
	}

	second, _ := homedir.Dir()
	second += "/" + file
	if FileExists(second) {
		return second
	}

	third := "/etc/" + file
	if FileExists(third) {
		return third
	}

	return file

}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
