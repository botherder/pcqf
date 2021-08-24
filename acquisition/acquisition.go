// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
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
	StoragePath      string    `json:"storage_path"`
	AutorunsExesPath string    `json:"autoruns_exes_path"`
	ProcsExesPath    string    `json:"procs_exes_path"`
	MemoryPath       string    `json:"memory_path"`
}

func New() (*Acquisition, error) {
	acq := Acquisition{
		UUID:     uuid.NewV4().String(),
		Datetime: time.Now().UTC(),
	}

	acq.StoragePath = filepath.Join(utils.GetCwd(), acq.UUID)
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
