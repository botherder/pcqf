package main

import (
	"fmt"
	"github.com/matishsiao/goInfo"
	"os"
	"os/user"
	"path"
)

// Get current working directory.
func getCwd() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}

	return path.Dir(exe)
}

// Get current username.
func getUserName() string {
	userObject, err := user.Current()
	if err != nil {
		return ""
	}

	return userObject.Username
}

// Get computer name.
func getComputerName() string {
	hostname, _ := os.Hostname()
	return hostname
}

// Get some accurate version of the operating system.
func getOperatingSystem() string {
	gi := goInfo.GetInfo()
	return fmt.Sprintf("%s %s %s", gi.OS, gi.Core, gi.Platform)
}
