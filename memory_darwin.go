package main

import (
	"os/exec"
	"io/ioutil"
	"path/filepath"
	log "github.com/Sirupsen/logrus"
)

osxpmemPath := filepath.Join(binPath, "osxpmem.app", "osxpmem")

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

	err = exec.Command("unzip", "-d", binpath).Run()
	if err != nil {
		return err
	}
}

func generateMemoryDump() {
	log.Info("Taking a snapshot of the system memory...")

	err := dropOSXPmem()
	if err != nil {
		log.Error("Unable to find OSXPmem: ", err.Error())
		return
	}
}
