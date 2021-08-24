// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/botherder/pcqf/utils"
	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/v3/mem"
)

var binPath = filepath.Join(utils.GetCwd(), "bin")

func initBinFolder() error {
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		err = os.MkdirAll(binPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func promptMemory() bool {
	virt, _ := mem.VirtualMemory()
	virtTotal := virt.Total / (1000 * 1000 * 1000)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Take a memory snapshot (it would take circa %d GB of space)", virtTotal),
		IsConfirm: true,
	}

	takeMemory, err := prompt.Run()
	if err != nil {
		printError("Failed to get choice for memory prompt", err)
		return false
	}

	if takeMemory == "y" {
		return true
	} else {
		return false
	}
}
