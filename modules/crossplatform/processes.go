// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package crossplatform

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/botherder/go-savetime/files"
	"github.com/botherder/pcqf/acquisition"
	"github.com/shirou/gopsutil/v3/process"
)

// RunningProcess defines the relevant details about a process.
type RunningProcess struct {
	Pid       int32  `json:"pid"`
	Name      string `json:"name"`
	ParentPid int32  `json:"ppid"`
	Exe       string `json:"exe"`
	Cmdline   string `json:"cmd"`
}

func GetProcessList(acq *acquisition.Acquisition) error {
	fmt.Println("Generating list of running processes...")

	procs, err := process.Processes()
	if err != nil {
		return fmt.Errorf("failed to create process list: %v", err)
	}

	var procsList []RunningProcess
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
		procsList = append(procsList, entry)

		if _, err := os.Stat(entry.Exe); err == nil {
			copyName := fmt.Sprintf("%d_%s.bin", entry.Pid, entry.Name)
			copyPath := filepath.Join(acq.ProcsExesPath, copyName)
			files.Copy(entry.Exe, copyPath)
		}
	}

	procsListPath := filepath.Join(acq.StoragePath, "process_list.json")
	procsListJSON, err := os.Create(procsListPath)
	if err != nil {
		return fmt.Errorf("failed to save process list to file %s: %v",
			procsListPath, err)
	}
	defer procsListJSON.Close()

	buf, _ := json.MarshalIndent(procsList, "", "    ")

	procsListJSON.WriteString(string(buf[:]))
	procsListJSON.Sync()

	return nil
}
