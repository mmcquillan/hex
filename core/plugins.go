package core

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/mmcquillan/hex-plugin"
	"github.com/mmcquillan/hex/models"
)

func Plugins(plugins *map[string]models.Plugin, rules map[string]models.Rule, config models.Config) {

	// look for existing plugins
	if config.PluginsDir != "" {
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

	// review rules for missing plugins
	for _, rule := range rules {
		for _, action := range rule.Actions {
			if !checkPlugin(action.Type, *plugins) {
				getPlugin(action.Type, config)
				addPlugin(action.Type, MakePath(config.PluginsDir, action.Type), *plugins, config)
			}
		}
	}

}

func StopPlugins(plugins map[string]models.Plugin, config models.Config) {
	for pluginName, _ := range plugins {
		config.Logger.Debug("Stopping Plugin " + pluginName)
		plugins[pluginName].Client.Kill()
	}
}

func checkPlugin(pluginName string, plugins map[string]models.Plugin) bool {
	if _, exists := plugins[pluginName]; exists {
		return true
	}
	return false
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

func getPlugin(pluginName string, config models.Config) {

	// check for exisitng file and remove if exists
	filePath := MakePath(config.PluginsDir, pluginName)
	if FileExists(filePath) {
		config.Logger.Debug("Plugin file being overwritten - " + filePath)
		os.Remove(filePath)
	}

	// generate a URL
	url := "https://hexbot.io/downloads/" + config.Version + "/" + pluginName

	// create new file for output
	output, err := os.Create(filePath)
	if err != nil {
		config.Logger.Error("Plugin create error - " + filePath + " " + err.Error())
		return
	}
	defer output.Close()

	// download the file
	response, err := http.Get(url)
	if err != nil {
		config.Logger.Error("Plugin download error - " + url + " " + err.Error())
		return
	}
	defer response.Body.Close()
	n, err := io.Copy(output, response.Body)
	if err != nil {
		config.Logger.Error("Plugin write error - " + filePath + " " + err.Error())
		return
	}
	err = os.Chmod(filePath, os.ModePerm)
	if err != nil {
		config.Logger.Error("Plugin Chmod error - " + filePath + " " + err.Error())
	}
	config.Logger.Debug("Downloaded " + strconv.FormatInt(n, 10) + " bytes for plugin " + pluginName)

	// download the md5 for file
	responseMd5, err := http.Get(url + ".md5")
	if err != nil {
		config.Logger.Error("Plugin download error - " + url + " " + err.Error())
		return
	}
	defer responseMd5.Body.Close()
	rawbody, err := ioutil.ReadAll(responseMd5.Body)
	if err != nil {
		config.Logger.Error("Plugin MD5 Read - " + err.Error())
	}
	md5chk := strings.TrimSpace(string(rawbody))

	// determine md5 of downloaded file
	fileMd5, err := os.Open(filePath)
	if err != nil {
		config.Logger.Error("Plugin MD5 File Read - " + err.Error())
	}
	defer fileMd5.Close()
	hash := md5.New()
	_, err = io.Copy(hash, fileMd5)
	if err != nil {
		config.Logger.Error("Plugin MD5 Eval - " + err.Error())
	}
	fileMd5chk := fmt.Sprintf("%x", hash.Sum(nil))

	// evaluate md5 match
	if fileMd5chk != md5chk {
		config.Logger.Error("MD5 Eval Failed for " + filePath + " [ " + string(hash.Sum(nil)) + " | " + md5chk + " ]")
		os.Remove(filePath)
	} else {
		config.Logger.Info("Downloaded Plugin " + filePath)
	}

}

func ResolvePluginName(pluginName string) string {
	subs := map[string]string{
		"local":    "hex-local",
		"response": "hex-response",
		"ssh":      "hex-ssh",
		"twilio":   "hex-twilio",
		"validate": "hex-validate",
		"winrm":    "hex-winrm",
	}
	if match, exists := subs[pluginName]; exists {
		pluginName = match
	}
	return pluginName
}
