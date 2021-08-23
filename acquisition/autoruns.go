// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/botherder/go-autoruns"
	"github.com/botherder/go-savetime/files"
)

func (a *Acquisition) GenerateAutoruns() error {
	// Fetch autoruns.
	autoruns := autoruns.Autoruns()

	// Make backup of autoruns executables.
	for _, autorun := range autoruns {
		if _, err := os.Stat(autorun.ImagePath); err == nil {
			copyName := fmt.Sprintf("%s_%s.bin", autorun.MD5, autorun.ImageName)
			copyPath := filepath.Join(a.AutorunsExesPath, copyName)
			files.Copy(autorun.ImagePath, copyPath)
		}
	}

	// Store the json list to file.
	autorunsJSONPath := filepath.Join(a.StoragePath, "autoruns.json")
	autorunsJSON, err := os.Create(autorunsJSONPath)
	if err != nil {
		return fmt.Errorf("failed to save autoruns to file: %v", err)
	}
	defer autorunsJSON.Close()

	// Encoding into json.
	buf, _ := json.MarshalIndent(autoruns, "", "    ")

	autorunsJSON.WriteString(string(buf[:]))
	autorunsJSON.Sync()

	return nil
}
