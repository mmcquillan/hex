package core

import (
	"fmt"
	"os"
	"time"

	"github.com/hexbotio/hex/models"
	"gopkg.in/src-d/go-git.v4"
)

func Pull(config models.Config) {
	if config.RulesGitUrl != "" && config.RulesDir != "" {
		if DirExists(config.RulesDir) {
			go watchGit(config)
		} else {
			fmt.Println("ERROR: The rules directory does not exist.")
			os.Exit(1)
		}
	}
}

func watchGit(config models.Config) {
	config.Logger.Trace("Pulling Rules Repo")
	_, err := git.PlainClone(config.RulesDir, false, &git.CloneOptions{
		URL:      config.RulesGitUrl,
		Progress: nil,
	})
	if err != nil {
		config.Logger.Error("Pull Error: " + err.Error())
	}
	repo, err := git.PlainOpen(config.RulesDir)
	if err != nil {
		config.Logger.Error("Pull Repo Error: " + err.Error())
	}
	worktree, err := repo.Worktree()
	if err != nil {
		if err.Error() == "repository already exists" {
			config.Logger.Trace("Rules Repo already exists")
		} else {
			config.Logger.Error("Pull Worktree Error: " + err.Error())
		}
	} else {
		config.Logger.Trace("Rules Repo pulled")
	}
	for {
		time.Sleep(30 * time.Second)
		config.Logger.Trace("Refreshing Rules Repo")
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			if err.Error() == "already up-to-date" {
				config.Logger.Trace("Rules Repo already up to date")
			} else {
				config.Logger.Error("Pull Refresh Error: " + err.Error())
			}
		} else {
			config.Logger.Trace("Rules Repo updated")
		}
	}
}
