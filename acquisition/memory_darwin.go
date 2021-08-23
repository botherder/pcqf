// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func dropOSXPmem() error {
	err := initBinFolder()
	if err != nil {
		return err
	}

	osxpmemData, err := Asset("osxpmem.zip")
	if err != nil {
		return err
	}

	zipPath := filepath.Join(binPath, "osxpmem.zip")
	err = ioutil.WriteFile(zipPath, osxpmemData, 0644)
	if err != nil {
		return err
	}

	err = exec.Command("unzip", zipPath, "-d", binPath).Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *Acquisition) GenerateMemoryDump() error {
	err := dropOSXPmem()
	if err != nil {
		return fmt.Errorf("failed to create osxpmem: %v", err)
	}

	osxpmemPath := filepath.Join(binPath, "osxpmem.app", "osxpmem")
	if _, err := os.Stat(osxpmemPath); os.IsNotExist(err) {
		return fmt.Errorf("failed to find osxpmem at path %s: %v",
			osxpmemPath, err)
	}

	cmdArgs := []string{"--format", "raw", "--output", a.MemoryPath}

	err = exec.Command(osxpmemPath, cmdArgs...).Run()
	if err != nil {
		return fmt.Errorf("failed to launch osxpmem (did you launch this with sudo?): %v",
			err)
	}

	return nil
}
