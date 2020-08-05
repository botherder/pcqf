// Copyright (c) 2017-2020 Claudio Guarnieri.
//
// This file is part of Snoopdigg.
//
// Snoopdigg is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Snoopdigg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Snoopdigg.  If not, see <https://www.gnu.org/licenses/>.

package main

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
		log.Error("Unable to launch OSXPmem (did you launch this with sudo?): ", err.Error())
		return
	}

	log.Info("Memory dump generated at ", acq.Memory)
}
