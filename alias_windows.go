// +build windows

package main

import "os/exec"

func addAlias(aliasName string, command string) (err error) {
	cmd := exec.Command("doskey", aliasName+"="+command)
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}
