// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/botherder/pcqf/utils"
	"github.com/satori/go.uuid"
)

// Acquisition defines the basic properties we want to store.
type Acquisition struct {
	UUID         string `json:"uuid"`
	Datetime     time.Time `json:"datetime"`
	ComputerName string `json:"computer_name"`
	ComputerUser string `json:"computer_user"`
	Platform     string `json:"platform"`
	Folder       string `json:"folder"`
	StoragePath      string `json:"storage"`
	AutorunsExesPath string `json:"autoruns_exes"`
	ProcsExesPath    string `json:"procs_exes"`
	MemoryPath       string `json:"memory"`
}

func New() (*Acquisition, error) {
	acq := Acquisition{
		UUID:         uuid.NewV4().String(),
		Datetime:     time.Now().UTC()
		ComputerName: utils.GetComputerName(),
		ComputerUser: utils.GetUserName(),
		Platform:     utils.GetOperatingSystem(),
	}

	// This is some spaghetti code to generate the folder name for the current
	// acquisition. It will just try to append a number until it finds a
	// combination that has not been used yet.
	baseFolder := fmt.Sprintf("%s_%s", acq.Date, acq.ComputerName)
	baseStorage := filepath.Join(utils.GetCwd())
	tmpFolder := baseFolder
	tmpStorage := filepath.Join(baseStorage, baseFolder)
	counter := 1
	for {
		if _, err := os.Stat(tmpStorage); os.IsNotExist(err) {
			break
		}

		tmpFolder = fmt.Sprintf("%s_%d", baseFolder, counter)
		tmpStorage = filepath.Join(baseStorage, tmpFolder)
		counter++
	}

	// Proceeds creating all the required subfolders.
	acq.Folder = tmpFolder
	acq.StoragePath = tmpStorage
	acq.AutorunsExesPath = filepath.Join(acq.StoragePath, "autoruns_bins")
	acq.ProcsExesPath = filepath.Join(acq.StoragePath, "process_bins")
	acq.MemoryPath = filepath.Join(acq.StoragePath, "memory")

	err := acq.createFolders()
	if err != nil {
		return nil, err
	}

	return &acq, nil
}

func (a *Acquisition) createFolders() error {
	folders := []string{a.AutorunsExesPath, a.ProcsExesPath, a.Memory}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
