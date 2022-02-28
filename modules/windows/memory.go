// pcqf - PC Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package windows

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/botherder/pcqf/acquisition"
	"github.com/botherder/pcqf/utils"
	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/v3/mem"
)

func dropWinpmem() error {
	winpmemData, err := Asset("winpmem.exe")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(utils.GetCwd(), "winpmem.exe"),
		winpmemData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func promptMemory() (bool, error) {
	virt, _ := mem.VirtualMemory()
	virtTotal := virt.Total / (1000 * 1000 * 1000)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Take a memory snapshot (it would take circa %d GB of space)", virtTotal),
		IsConfirm: true,
	}

	takeMemory, err := prompt.Run()
	if err != nil {
		return false, err
	}

	if takeMemory == "y" {
		return true, nil
	} else {
		return false, nil
	}
}

func GenerateMemoryDump(acq *acquisition.Acquisition) error {
	takeMemory, err := promptMemory()
	if err != nil {
		return fmt.Errorf("failed to get choice on memory snapshot: %v",
			err)
	}
	if !takeMemory {
		fmt.Println("Skipping memory acquisition.")
		return nil
	}

	err = dropWinpmem()
	if err != nil {
		return fmt.Errorf("failed to create winpmem: %v", err)
	}

	cmdArgs := []string{filepath.Join(acq.MemoryPath, "physmem.raw")}
	err = exec.Command(filepath.Join(utils.GetCwd(), "winpmem.exe"),
		cmdArgs...).Run()
	if err != nil {
		return fmt.Errorf("failed to launch winpmem (are you running as Administrator?): %v", err)
	}

	return nil
}
