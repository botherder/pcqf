package main

import (
	"os"
	"os/exec"
	"io/ioutil"
	"path/filepath"
	log "github.com/Sirupsen/logrus"
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

func generateMemoryDump() {
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

	cmdArgs := []string{"--format", "raw", "--output", acq.Memory}

	err = exec.Command(osxpmemPath, cmdArgs...).Run()
	if err != nil {
		log.Error("Unable to launch OSXPmem (did you launch this with sudo?)")
		return
	}

	log.Info("Memory dump generated at ", acq.Memory)
}
