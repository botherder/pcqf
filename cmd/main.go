// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"fmt"
	"os"

	"github.com/botherder/pcqf/acquisition"
	"github.com/manifoldco/promptui"
	"github.com/mattn/go-colorable"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

var acq acquisition.Acquisition

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
}

func main() {
	acq, err := acquisition.New()
	if err != nil {
		log.Error(err)
		return
	}

	err = acq.CreateFolders()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Started acquisition ", acq.Folder)

	generateProfile()
	generateProcessList()
	generateAutoruns()

	virt, _ := mem.VirtualMemory()
	virtTotal := virt.Total / (1000 * 1000 * 1000)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Take a memory snapshot (it would take circa %d GB of space)", virtTotal),
		IsConfirm: true,
	}

	takeMemory, err := prompt.Run()
	if err == nil && takeMemory == "y" {
		generateMemoryDump()
	} else {
		log.Info("Skipping memory acquisition.")
	}

	storeSecurely()

	log.Info("Acquisition completed.")

	log.Info("Press Enter to finish ...")
	var b = make([]byte, 1)
	os.Stdin.Read(b)
}
