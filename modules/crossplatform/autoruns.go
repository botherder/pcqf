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

	"github.com/botherder/go-autoruns"
	"github.com/botherder/go-savetime/files"
	"github.com/botherder/pcqf/acquisition"
)

func GetAutoruns(acq *acquisition.Acquisition) error {
	fmt.Println("Generating list of persistent software...")

	autoruns := autoruns.Autoruns()

	// Make backup of autoruns executables.
	for _, autorun := range autoruns {
		if _, err := os.Stat(autorun.ImagePath); err == nil {
			copyName := fmt.Sprintf("%s_%s.bin", autorun.MD5, autorun.ImageName)
			copyPath := filepath.Join(acq.AutorunsExesPath, copyName)
			files.Copy(autorun.ImagePath, copyPath)
		}
	}

	autorunsJSONPath := filepath.Join(acq.StoragePath, "autoruns.json")
	autorunsJSON, err := os.Create(autorunsJSONPath)
	if err != nil {
		return fmt.Errorf("failed to save autoruns to file: %v", err)
	}
	defer autorunsJSON.Close()

	buf, _ := json.MarshalIndent(autoruns, "", "    ")
	autorunsJSON.WriteString(string(buf[:]))
	autorunsJSON.Sync()

	return nil
}
