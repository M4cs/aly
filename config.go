package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

func checkForConfig() (config Config, err error) {
	usr, err := user.Current()
	if err != nil {
		return config, err
	}
	if _, err := os.Stat(path.Join(usr.HomeDir + "/.alyconfig.json")); os.IsNotExist(err) {
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

func (config *Config) updateJSON() (err error) {
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

func (config *Config) addPlugin(plugin Plugin) {
	config.EnabledPlugins = append(config.EnabledPlugins, plugin)
	config.updateJSON()
}

func (config *Config) enablePlugin(name string) (err error) {
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			config.DisabledPlugins = append(config.DisabledPlugins[:i], config.DisabledPlugins[i+1:]...)
			config.EnabledPlugins = append(config.EnabledPlugins, plugin)
			return nil
		}
	}
	return errors.New("couldn't find a plugin with that name")
}

func (config *Config) disablePlugin(name string) (err error) {
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			config.EnabledPlugins = append(config.EnabledPlugins[:i], config.EnabledPlugins[i+1:]...)
			config.DisabledPlugins = append(config.DisabledPlugins, plugin)
			return nil
		}
	}
	return errors.New("couldn't find a plugin with that name")
}
