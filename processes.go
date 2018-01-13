package main

import (
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/shirou/gopsutil/process"
	log "github.com/Sirupsen/logrus"
)

type RunningProcess struct {
	Pid int32 `json:"pid"`
	Name string `json:"name"`
	ParentPid int32 `json:"ppid"`
	Exe string `json:"exe"`
	Cmdline string `json:"cmd"`
}

func generateProcessList() {
	log.Info("Generating list of running processes...")

	procs, err := process.Processes()
	if err != nil {
		log.Error("Unable to create process list: ", err.Error())
		return
	}

	var processList []*RunningProcess
	for _, proc := range procs {
		entry := RunningProcess{
			Pid: proc.Pid(),
			Name: proc.Name(),
			ParentPid: proc.Ppid(),
			Exe: proc.Exe(),
			Cmdline: proc.Cmdline()
		}
		processList = append(processList, entry)
	}

	processListPath := filepath.Join(acq.Storage, "processlist.json")
	processListJson, err := os.Create(processListPath)
	if err != nil {
		log.Error("Unable to save process list to file: ", err.Error())
		return
	}
	defer processListJson.Close()

	buf, _ := json.MarshalIndent(processList, "", "    ")

	processListJson.WriteString(string(buf[:]))
	processListJson.Sync()

	log.Info("Process list generated successfully!")
}
