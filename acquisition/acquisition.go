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
	UUID         string
	Date         string
	Time         string
	ComputerName string
	ComputerUser string
	Platform     string
	Folder       string
	Storage      string
	AutorunsExes string
	ProcsExes    string
	Memory       string
}

func New() (*Acquisition, error) {
	acq := Acquisition{
		UUID:         uuid.NewV4().String(),
		Date:         time.Now().UTC().Format("2006-02-01"),
		Time:         time.Now().UTC().Format("15:04:05"),
		ComputerName: utils.GetComputerName(),
		ComputerUser: utils.GetUserName(),
		Platform:     utils.GetOperatingSystem(),
	}

	// This is some spaghetti code to generate the folder name for the current
	// acquisition. It will just try to append a number until it finds a
	// combination that has not been used yet.
	baseFolder := fmt.Sprintf("%s_%s", acq.Date, acq.ComputerName)
	baseStorage := filepath.Join(utils.GetCwd(), "acquisitions")
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
	acq.Storage = tmpStorage
	acq.AutorunsExes = filepath.Join(acq.Storage, "autoruns_bins")
	acq.ProcsExes = filepath.Join(acq.Storage, "process_bins")
	acq.Memory = filepath.Join(acq.Storage, "memory")

	return &acq, nil
}

func (a *Acquisition) CreateFolders() error {
	folders := []string{a.AutorunsExes, a.ProcsExes, a.Memory}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
