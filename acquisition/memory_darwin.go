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

	log "github.com/sirupsen/logrus"
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

func (a *Acquisition) GenerateMemoryDump() {
	log.Info("Taking a snapshot of the system memory...")

	err := dropOSXPmem()
	if err != nil {
		log.Error("Unable to find OSXPmem: ", err.Error())
		return
	}

	osxpmemPath := filepath.Join(binPath, "osxpmem.app", "osxpmem")
	if _, err := os.Stat(osxpmemPath); os.IsNotExist(err) {
		log.Error("Unable to find OSXPmem at path: ", osxpmemPath)
		return
	}

	cmdArgs := []string{"--format", "raw", "--output", a.Memory}

	err = exec.Command(osxpmemPath, cmdArgs...).Run()
	if err != nil {
		log.Error("Unable to launch OSXPmem (did you launch this with sudo?): ", err.Error())
		return
	}

	log.Info("Memory dump generated at ", a.Memory)
}
