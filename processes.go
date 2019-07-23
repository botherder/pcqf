// Copyright (c) 2017-2019 Claudio Guarnieri.
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
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/shirou/gopsutil/process"
	"os"
	"path/filepath"
)

type RunningProcess struct {
	Pid       int32  `json:"pid"`
	Name      string `json:"name"`
	ParentPid int32  `json:"ppid"`
	Exe       string `json:"exe"`
	Cmdline   string `json:"cmd"`
}

func generateProcessList() {
	log.Info("Generating list of running processes...")

	procs, err := process.Processes()
	if err != nil {
		log.Error("Unable to create process list: ", err.Error())
		return
	}

	var processList []RunningProcess
	for _, proc := range procs {
		name, _ := proc.Name()
		ppid, _ := proc.Ppid()
		exe, _ := proc.Exe()
		cmd, _ := proc.Cmdline()

		entry := RunningProcess{
			Pid:       proc.Pid,
			Name:      name,
			ParentPid: ppid,
			Exe:       exe,
			Cmdline:   cmd,
		}
		processList = append(processList, entry)
	}

	processListPath := filepath.Join(acq.Storage, "processlist.json")
	processListJSON, err := os.Create(processListPath)
	if err != nil {
		log.Error("Unable to save process list to file: ", err.Error())
		return
	}
	defer processListJSON.Close()

	buf, _ := json.MarshalIndent(processList, "", "    ")

	processListJSON.WriteString(string(buf[:]))
	processListJSON.Sync()

	log.Info("Process list generated successfully!")
}
