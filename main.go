package main

import (
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("aly", "Alias manager for all platforms")
	var load *bool = parser.Flag("l", "load", &argparse.Options{Required: false, Help: "Load aliases into shell session. Should be added to profile", Default: true})
	var localFile *bool = parser.Flag("f", "local-file", &argparse.Options{Required: false, Help: "Pass to add plugin from local file", Default: false})
	add := parser.String("a", "add", &argparse.Options{Required: false, Help: "Add a plugin by URL. Pass '-f' to use local file."})
	remove := parser.String("r", "remove", &argparse.Options{Required: false, Help: "Remove a plugin by name"})
	enable := parser.String("e", "enable", &argparse.Options{Required: false, Help: "Enable a plugin by name"})
	disable := parser.String("d", "disable", &argparse.Options{Required: false, Help: "Disable a plugin by name"})
	update := parser.String("u", "update", &argparse.Options{Required: false, Help: "Update a plugin by name. Requires remote URL in plugin"})
	var updateall *bool = parser.Flag("t", "total-update", &argparse.Options{Required: false, Help: "Update all plugins installed", Default: false})
	customConfig := parser.String("c", "config", &argparse.Options{Required: false, Help: "Config file to load from"})
	var info *bool = parser.Flag("i", "info", &argparse.Options{Required: false, Help: "See loaded aliases", Default: false})
	pluginInfo := parser.String("p", "plugin-info", &argparse.Options{Required: false, Help: "See plugin status and information"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fileName := ""
	if *customConfig != "" {
		fileName = *customConfig
	}
	config, err := checkForConfig(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if *updateall {
		config.updateAllPlugins()
	}
	if *update != "" {
		config.updatePlugin(*update)
	}
	if *remove != "" {
		config.removePlugin(*remove)
	}
	if *info {
		config.alyStatus()
	}
	if *pluginInfo != "" {
		config.pluginInfo(*pluginInfo)
	}
	if *add != "" {
		config.addPlugin(*add, *localFile)
	}
	if *enable != "" {
		config.enablePlugin(*enable)
	}
	if *disable != "" {
		config.disablePlugin(*disable)
	}
	if *load {
		config.loadPlugins()
	}
}
