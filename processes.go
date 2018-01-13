package main

import (
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/shirou/gopsutil/process"
	log "github.com/Sirupsen/logrus"
)

func generateProcessList() {
	log.Info("Generating list of running processes...")

	processes, err := process.Processes()
	if err != nil {
		log.Error("Unable to create process list: ", err.Error())
		return
	}

	processesJsonPath := filepath.Join(acq.Storage, "processes.json")
	processesJson, err := os.Create(processesJsonPath)
	if err != nil {
		log.Error("Unable to save process list to file: ", err.Error())
		return
	}
	defer processesJson.Close()

	buf, _ := json.MarshalIndent(processes, "", "    ")

	processesJson.WriteString(string(buf[:]))
	processesJson.Sync()

	log.Info("Process list generated successfully!")
}
