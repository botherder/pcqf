// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package utils

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/matishsiao/goInfo"
)

// Get current working directory.
func GetCwd() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}

	return path.Dir(exe)
}

// Get current username.
func GetUserName() string {
	userObject, err := user.Current()
	if err != nil {
		return ""
	}

	return userObject.Username
}

// Get computer name.
func GetComputerName() string {
	hostname, _ := os.Hostname()
	return hostname
}

// Get some accurate version of the operating system.
func GetOperatingSystem() string {
	gi := goInfo.GetInfo()
	return fmt.Sprintf("%s %s %s", gi.OS, gi.Core, gi.Platform)
}
