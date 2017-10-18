package core

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
	"github.com/hexbotio/hex/models"
)

func Plugins(plugins *map[string]models.Plugin, config models.Config) {
	if DirExists(config.PluginsDir) {
		pluginList := make(map[string]string)
		err := filepath.Walk(config.PluginsDir, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				pluginList[f.Name()] = path
			}
			return nil
		})
		if err != nil {
			config.Logger.Error("Loading Plugin Dir" + " - " + err.Error())
		}
		for pluginName, pluginPath := range pluginList {
			addPlugin(pluginName, pluginPath, *plugins, config)
		}
	} else {
		config.Logger.Error("Plugin Directory does not exist")
	}
}

func StopPlugins(plugins map[string]models.Plugin, config models.Config) {
	for pluginName, _ := range plugins {
		config.Logger.Debug("Stopping Plugin " + pluginName)
		plugins[pluginName].Client.Kill()
	}
}

func addPlugin(pluginName string, pluginPath string, plugins map[string]models.Plugin, config models.Config) {
	if _, exists := plugins[pluginName]; !exists {
		config.Logger.Info("Initializing Plugin " + pluginName + " (" + pluginPath + ")")
		if FileExists(pluginPath) {
			var pluginMap = map[string]plugin.Plugin{
				"action": &hexplugin.HexPlugin{},
			}
			client := plugin.NewClient(&plugin.ClientConfig{
				HandshakeConfig: hexplugin.GetHandshakeConfig(),
				Plugins:         pluginMap,
				Cmd:             exec.Command(pluginPath),
				Logger:          config.Logger,
			})
			rpcClient, err := client.Client()
			if err != nil {
				config.Logger.Error("RPC Client" + " - " + err.Error())
			}
			raw, err := rpcClient.Dispense("action")
			if err != nil {
				config.Logger.Error("Plugin Request" + " - " + err.Error())
			}
			plugins[pluginName] = models.Plugin{
				Name:   pluginName,
				Path:   pluginPath,
				Client: client,
				Action: raw.(hexplugin.Action),
			}
		} else {
			config.Logger.Error("Plugin Path Missing " + pluginPath)
		}
	}
}

func ResolvePlugin(plugin string) string {
	subs := map[string]string{
		"local":    "hex-local",
		"response": "hex-response",
		"ssh":      "hex-ssh",
	}
	if match, exists := subs[plugin]; exists {
		return match
	}
	return plugin
}
