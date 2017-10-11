package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/hexbotio/hex/models"
)

var fileFilter = ".json"

func Rules(rules *map[string]models.Rule, config models.Config) {
	if DirExists(config.RulesDir) {
		go watchRules(config, rules)
		ruleList := []string{}
		err := filepath.Walk(config.RulesDir, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && strings.HasSuffix(f.Name(), fileFilter) {
				ruleList = append(ruleList, path)
			}
			return nil
		})
		if err != nil {
			config.Logger.Error("Loading Rules Directory", err)
		}
		for _, file := range ruleList {
			addRule(file, *rules, config)
		}
	} else {
		fmt.Println("ERROR: The rules directory does not exist.")
		os.Exit(1)
	}
}

func addRule(ruleFile string, rules map[string]models.Rule, config models.Config) {
	if _, exists := rules[ruleFile]; !exists {
		if strings.HasSuffix(ruleFile, fileFilter) {
			config.Logger.Info("Loading Rule " + ruleFile)
			rules[ruleFile] = readRule(ruleFile, config)
		}
	}
}

func reloadRule(ruleFile string, rules map[string]models.Rule, config models.Config) {
	if strings.HasSuffix(ruleFile, fileFilter) {
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
		Active: true,
		Debug:  false,
		Hide:   false,
		ACL:    "*",
	}
	if FileExists(ruleFile) {
		file, err := ioutil.ReadFile(ruleFile)
		if err != nil {
			config.Logger.Error("Add Rule File Read "+ruleFile, err)
			rule.Active = false
		}
		err = json.Unmarshal(file, &rule)
		if err != nil {
			config.Logger.Error("Add Rule Unmarshal "+ruleFile, err)
			rule.Active = false
		}
	}
	return rule
}

func watchRules(config models.Config, rules *map[string]models.Rule) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		config.Logger.Error("File Watcher", err)
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
				config.Logger.Error("Rule Load", err)
			}
		}
	}()

	err = watcher.Add(config.RulesDir)
	if err != nil {
		config.Logger.Error("File Watcher Add", err)
	}
	<-done

}
