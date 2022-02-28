// pcqf - PC Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package crossplatform

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/botherder/pcqf/acquisition"
	"github.com/matishsiao/goInfo"
)

type Profile struct {
	ComputerName string `json:"computer_name"`
	ComputerUser string `json:"computer_user"`
	Platform     string `json:"platform"`
}

func getUserName() string {
	userObject, err := user.Current()
	if err != nil {
		return ""
	}

	return userObject.Username
}

func getComputerName() string {
	hostname, _ := os.Hostname()
	return hostname
}

func getOperatingSystem() string {
	gi := goInfo.GetInfo()
	return fmt.Sprintf("%s %s %s", gi.OS, gi.Core, gi.Platform)
}

func GenerateSystemInfo(acq *acquisition.Acquisition) error {
	fmt.Println("Generating system info...")

	profilePath := filepath.Join(acq.StoragePath, "systeminfo.json")
	profileFile, err := os.Create(profilePath)
	if err != nil {
		return fmt.Errorf("failed to generate systeminfo file: %v", err)
	}
	defer profileFile.Close()

	profile := Profile{
		ComputerName: getComputerName(),
		ComputerUser: getUserName(),
		Platform:     getOperatingSystem(),
	}

	buf, _ := json.MarshalIndent(profile, "", "    ")

	profileFile.WriteString(string(buf[:]))
	profileFile.Sync()

	return nil
}
