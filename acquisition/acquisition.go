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
	UUID             string    `json:"uuid"`
	Datetime         time.Time `json:"datetime"`
	ComputerName     string    `json:"computer_name"`
	ComputerUser     string    `json:"computer_user"`
	Platform         string    `json:"platform"`
	FolderName       string    `json:"folder_name"`
	StoragePath      string    `json:"storage_path"`
	AutorunsExesPath string    `json:"autoruns_exes_path"`
	ProcsExesPath    string    `json:"procs_exes_path"`
	MemoryPath       string    `json:"memory_path"`
}

func New() (*Acquisition, error) {
	acq := Acquisition{
		UUID:         uuid.NewV4().String(),
		Datetime:     time.Now().UTC(),
		ComputerName: utils.GetComputerName(),
		ComputerUser: utils.GetUserName(),
		Platform:     utils.GetOperatingSystem(),
	}

	// This is some spaghetti code to generate the folder name for the current
	// acquisition. It will just try to append a number until it finds a
	// combination that has not been used yet.
	cwd := utils.GetCwd()

	tmpFolderName := acq.UUID
	tmpStoragePath := filepath.Join(cwd, acq.UUID)

	counter := 1
	for {
		// If the current tmpStoragePath does not exist, it is fine to use,
		// so we break out of the loop.
		if _, err := os.Stat(tmpStoragePath); os.IsNotExist(err) {
			break
		}

		// Otherwise, we try again appending the current counter to the
		// folder name.
		tmpFolderName = fmt.Sprintf("%s_%d", acq.UUID, counter)
		tmpStoragePath = filepath.Join(cwd, tmpFolderName)
		counter++
	}

	// Proceeds creating all the required subfolders.
	acq.FolderName = tmpFolderName
	acq.StoragePath = tmpStoragePath
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
	folders := []string{a.AutorunsExesPath, a.ProcsExesPath, a.MemoryPath}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
