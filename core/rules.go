package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/mmcquillan/hex/models"
	"gopkg.in/yaml.v2"
)

func Rules(rules *map[string]models.Rule, config models.Config) {
	if config.RulesDir != "" {
		if DirExists(config.RulesDir) {
			go watchRules(config, rules)
			ruleList := []string{}
			err := filepath.Walk(config.RulesDir, func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() && isConfig(f.Name()) {
					ruleList = append(ruleList, path)
				}
				return nil
			})
			if err != nil {
				config.Logger.Error("Loading Rules Directory" + " - " + err.Error())
			}
			for _, file := range ruleList {
				addRule(file, *rules, config)
			}
		} else {
			fmt.Println("ERROR: The rules directory does not exist.")
			os.Exit(1)
		}
	}
}

func addRule(ruleFile string, rules map[string]models.Rule, config models.Config) {
	if _, exists := rules[ruleFile]; !exists {
		if isConfig(ruleFile) {
			config.Logger.Info("Loading Rule " + ruleFile)
			rules[ruleFile] = readRule(ruleFile, config)
		}
	}
}

func reloadRule(ruleFile string, rules map[string]models.Rule, config models.Config) {
	if isConfig(ruleFile) {
		config.Logger.Info("Reloading Rule " + ruleFile)
		rules[ruleFile] = readRule(ruleFile, config)
	}
}

func removeRule(ruleFile string, rules map[string]models.Rule, config models.Config) {
	if _, chk := rules[ruleFile]; chk {
		config.Logger.Info("Removing Rule " + ruleFile)
		delete(rules, ruleFile)
	}
}

func readRule(ruleFile string, config models.Config) (rule models.Rule) {
	rule = models.Rule{
		Active:         true,
		Debug:          false,
		Format:         false,
		Hide:           false,
		ACL:            "*",
		OutputFailOnly: false,
		OutputOnChange: false,
	}
	if FileExists(ruleFile) {
		file, err := ioutil.ReadFile(ruleFile)
		if err != nil {
			config.Logger.Error("Add Rule File Read " + ruleFile + " - " + err.Error())
			rule.Active = false
			return rule
		}
		ruleType := fileType(ruleFile)
		if ruleType == "json" {
			err = json.Unmarshal(file, &rule)
			if err != nil {
				config.Logger.Error("Add Rule json Unmarshal " + ruleFile + " - " + err.Error())
				rule.Active = false
				return rule
			}
		} else if ruleType == "yaml" {
			err = yaml.Unmarshal(file, &rule)
			if err != nil {
				config.Logger.Error("Add Rule yaml Unmarshal " + ruleFile + " - " + err.Error())
				rule.Active = false
				return rule
			}
		} else {
			return rule
		}
		for i := 0; i < len(rule.Actions); i++ {
			rule.Actions[i].Type = ResolvePluginName(rule.Actions[i].Type)
		}
		// no need to sub action.config as this happens at matcher time
	}
	rule.Id = ruleFile
	return rule
}

func watchRules(config models.Config, rules *map[string]models.Rule) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		config.Logger.Error("File Watcher" + " - " + err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					addRule(event.Name, *rules, config)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					removeRule(event.Name, *rules, config)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					reloadRule(event.Name, *rules, config)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					removeRule(event.Name, *rules, config)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					reloadRule(event.Name, *rules, config)
				}
			case err := <-watcher.Errors:
				config.Logger.Error("Rule Load" + " - " + err.Error())
			}
		}
	}()

	err = watcher.Add(config.RulesDir)
	if err != nil {
		config.Logger.Error("File Watcher Add" + " - " + err.Error())
	}
	<-done

}

func isConfig(file string) bool {
	if strings.HasSuffix(file, ".json") {
		return true
	}
	if strings.HasSuffix(file, ".yaml") {
		return true
	}
	if strings.HasSuffix(file, ".yml") {
		return true
	}
	return false
}

func fileType(file string) string {
	if strings.HasSuffix(file, ".json") {
		return "json"
	}
	if strings.HasSuffix(file, ".yaml") {
		return "yaml"
	}
	if strings.HasSuffix(file, ".yml") {
		return "yaml"
	}
	return ""
}
