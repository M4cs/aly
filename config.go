package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

func checkForConfig(filename string) (config Config, err error) {
	usr, err := user.Current()
	if err != nil {
		return config, err
	}
	if filename == "" {
		filename = path.Join(usr.HomeDir + "/.alyconfig.json")
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		config = createConfig()
		config.updateJSON()
	} else {
		file, err := ioutil.ReadFile(path.Join(usr.HomeDir + "/.alyconfig.json"))
		if err != nil {
			return config, err
		}
		json.Unmarshal(file, &config)
	}
	return config, nil
}

func (config Config) updateJSON() (err error) {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	configJSON, err := json.Marshal(&config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(usr.HomeDir+"/.alyconfig.json"), configJSON, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (config Config) allPlugins() (plugins []Plugin) {
	plugins = append(config.DisabledPlugins, config.EnabledPlugins...)
	return plugins
}

func createConfig() (config Config) {
	config = Config{
		Version:         "1.1.0",
		EnabledPlugins:  []Plugin{},
		DisabledPlugins: []Plugin{},
	}
	return config
}
