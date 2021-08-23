// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

var winpmemPath string = filepath.Join(binPath, "winpmem.exe")

func dropWinpmem() error {
	err := initBinFolder()
	if err != nil {
		return err
	}

	winpmemData, err := Asset("winpmem.exe")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(winpmemPath, winpmemData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a *Acquisition) GenerateMemoryDump() error {
	err := dropWinpmem()
	if err != nil {
		return fmt.Errorf("failed to create winpmem: %v", err)
	}

	cmdArgs := []string{filepath.Join(a.MemoryPath, "physmem.raw")}

	err = exec.Command(winpmemPath, cmdArgs...).Run()
	if err != nil {
		return fmt.Errorf("failed to launch winpmem (are you running as Administrator?): %v", err)
	}

	return nil
}
