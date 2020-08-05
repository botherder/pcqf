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
	log "github.com/sirupsen/logrus"
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

func generateMemoryDump() {
	log.Info("Taking a snapshot of the system memory...")

	err := dropWinpmem()
	if err != nil {
		log.Error("Unable to find winpmem: ", err.Error())
		return
	}

	cmdArgs := []string{"--format", "raw", "--output", acq.Memory}

	err = exec.Command(winpmemPath, cmdArgs...).Run()
	if err != nil {
		log.Error("Unable to launch winpmem (are you running as Administrator?): ", err.Error())
		return
	}

	log.Info("Memory dump generated at ", acq.Memory)
}
