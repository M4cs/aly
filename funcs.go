package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

	"golang.org/x/mod/semver"
)

func addAlias(aliasName string, command string, f *os.File) (err error) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("doskey", aliasName+"="+command)
		if err = cmd.Run(); err != nil {
			return err
		}
	} else {
		if _, err := f.Write([]byte("alias " + aliasName + "='" + command + "'\n")); err != nil {
			return err
		}
	}
	return nil
}

func (config Config) clearEmpties() (err error) {
	if runtime.GOOS == "windows" {
		for _, plugin := range config.DisabledPlugins {
			for _, am := range plugin.AliasMap {
				parent := am.Name
				cmd := exec.Command("doskey", parent+"=")
				if err = cmd.Run(); err != nil {
					return err
				}
				if len(am.Subalias) > 0 {
					for saName := range am.Subalias {
						cmd := exec.Command("doskey", parent+saName+"=")
						if err = cmd.Run(); err != nil {
							return err
						}
					}
				}
			}

		}
	}
	return nil
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (config Config) updateAllPlugins() (err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	allPlugins := config.allPlugins()
	pluginsToUpdate := []Plugin{}
	for _, plugin := range allPlugins {
		if plugin.URL != "" {
			res, err := client.Get(plugin.URL)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			var pl Plugin
			json.NewDecoder(res.Body).Decode(&pl)
			if semver.Compare(plugin.Version, pl.Version) == -1 {
				fmt.Println("Updating " + plugin.Name + "| Version: " + plugin.Version + " -> " + pl.Version)
				pluginsToUpdate = append(pluginsToUpdate, pl)
			}
		}
	}
	if len(pluginsToUpdate) > 0 {
		for _, plugin := range pluginsToUpdate {
			for i, pl := range config.DisabledPlugins {
				if pl.Name == plugin.Name {
					config.DisabledPlugins[i] = plugin
					break
				}
			}
			for i, pl := range config.EnabledPlugins {
				if pl.Name == plugin.Name {
					config.EnabledPlugins[i] = plugin
					break
				}
			}
		}
		config.updateJSON()
	} else {
		fmt.Println("No Plugins To Update!")
	}
	return nil
}

func (config Config) updatePlugin(name string) (err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			res, err := client.Get(plugin.URL)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			var pl Plugin
			json.NewDecoder(res.Body).Decode(&pl)
			if (semver.Compare(plugin.Version, pl.Version)) == -1 {
				config.DisabledPlugins[i] = pl
				config.updateJSON()
				return nil
			}
		}
	}
	for i, plugin := range config.EnabledPlugins {
		if plugin.Name == name {
			res, err := client.Get(plugin.URL)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			var pl Plugin
			json.NewDecoder(res.Body).Decode(&pl)
			if (semver.Compare(plugin.Version, pl.Version)) == -1 {
				config.EnabledPlugins[i] = pl
				config.updateJSON()
				return nil
			}
		}
	}
	fmt.Println(name + " was not found as a plugin name! ChEcK yOuR cAsE!")
	return nil
}

func (config Config) removePlugin(name string) (err error) {
	for i, plugin := range config.EnabledPlugins {
		if plugin.Name == name {
			config.EnabledPlugins = append(config.EnabledPlugins[:i], config.EnabledPlugins[i+1:]...)
			config.updateJSON()
			fmt.Println("Removed " + name)
			return nil
		}
	}
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			config.EnabledPlugins = append(config.EnabledPlugins[:i], config.EnabledPlugins[i+1:]...)
			config.updateJSON()
			fmt.Println("Removed " + name)
			return nil
		}
	}
	fmt.Println(name + " was not found as a plugin! ChEcK cAsE!")
	return nil
}

func (config Config) enablePlugin(name string) (err error) {
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			config.DisabledPlugins = append(config.DisabledPlugins[:i], config.DisabledPlugins[i+1:]...)
			config.EnabledPlugins = append(config.EnabledPlugins, plugin)
			config.updateJSON()
			fmt.Println("Enabled Plugin: " + plugin.Name)
			return nil
		}
	}
	return errors.New("couldn't find a plugin with that name")
}

func (config Config) disablePlugin(name string) (err error) {
	for i, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			config.EnabledPlugins = append(config.EnabledPlugins[:i], config.EnabledPlugins[i+1:]...)
			config.DisabledPlugins = append(config.DisabledPlugins, plugin)
			config.updateJSON()
			fmt.Println("Disabled Plugin: " + plugin.Name)
			return nil
		}
	}
	return errors.New("couldn't find a plugin with that name")
}

func (config Config) addPlugin(url string, isLocal bool) (err error) {
	var plugin Plugin
	if isLocal {
		file, err := ioutil.ReadFile(url)
		if err != nil {
			return err
		}
		json.Unmarshal(file, &plugin)
	} else {
		client := &http.Client{Timeout: 10 * time.Second}
		res, err := client.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&plugin)
	}
	allPlugins := config.allPlugins()
	for _, p := range allPlugins {
		if p.Name == plugin.Name {
			return errors.New("you already have a plugin of that name, try updating it")
		}
	}
	config.DisabledPlugins = append(config.DisabledPlugins, plugin)
	config.updateJSON()
	fmt.Println("Added new plugin successfully! " + plugin.Name + " " + plugin.Version)
	return nil
}

func (config Config) alyStatus() (err error) {
	fmt.Println("Enabled Plugins:")
	for _, plugin := range config.EnabledPlugins {
		fmt.Println("Name: " + plugin.Name)
		fmt.Println("Version: " + plugin.Version)
		fmt.Println("Description: " + plugin.Description)
		fmt.Println("Author: " + plugin.Author)
		fmt.Println("URL: " + plugin.URL)
	}
	fmt.Println("\nDisabled Plugins:")
	for _, plugin := range config.DisabledPlugins {
		fmt.Println("Name: " + plugin.Name)
		fmt.Println("Version: " + plugin.Version)
		fmt.Println("Description: " + plugin.Description)
		fmt.Println("Author: " + plugin.Author)
		fmt.Println("URL: " + plugin.URL)
	}
	fmt.Println("\nUse '-p PLUGIN_NAME' to see status of certain plugin.")
	return nil
}

func (config Config) pluginInfo(name string) (err error) {
	for _, plugin := range config.EnabledPlugins {
		if plugin.Name == name {
			fmt.Println("Name: " + plugin.Name)
			fmt.Println("Version: " + plugin.Version)
			fmt.Println("Description: " + plugin.Description)
			fmt.Println("Author: " + plugin.Author)
			fmt.Println("URL: " + plugin.URL)
			fmt.Println("Enabled: TRUE")
		}
	}
	for _, plugin := range config.DisabledPlugins {
		if plugin.Name == name {
			fmt.Println("Name: " + plugin.Name)
			fmt.Println("Version: " + plugin.Version)
			fmt.Println("Description: " + plugin.Description)
			fmt.Println("Author: " + plugin.Author)
			fmt.Println("URL: " + plugin.URL)
			fmt.Println("Enabled: FALSE")
		}
	}
	return nil
}

func (config Config) loadPlugins() (err error) {
	usrHD, err := os.UserHomeDir()
	var f *os.File
	if err != nil {
		return err
	}
	if _, err := os.Stat(path.Join(usrHD, "/.alyases")); !os.IsNotExist(err) {
		os.Remove(path.Join(usrHD, "/.alyases"))
	}
	f, err = os.OpenFile(path.Join(usrHD, "/.alyases"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	config.clearEmpties()
	var loadedAliases []string
	for _, plugin := range config.EnabledPlugins {
		for _, alias := range plugin.AliasMap {
			parent := alias.Name
			_, found := find(loadedAliases, parent)
			if found {
				fmt.Println("[aly] Not loading " + plugin.Name + "... alias '" + parent + "' already found in another plugin!")
				continue
			}
			command := alias.Command
			err := addAlias(parent, command, f)
			loadedAliases = append(loadedAliases, parent)
			if err != nil {
				return err
			}
			if len(alias.Subalias) > 0 {
				for saAlias, saCommand := range alias.Subalias {
					newAlias := parent + saAlias
					_, found = find(loadedAliases, newAlias)
					if found {
						fmt.Println("[aly] Not loading " + plugin.Name + "'s subalias: '" + parent + "'. It was already found in another plugin!")
					} else {
						err := addAlias(newAlias, saCommand, f)
						if err != nil {
							return err
						}
						loadedAliases = append(loadedAliases, newAlias)
					}
				}
			}
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
