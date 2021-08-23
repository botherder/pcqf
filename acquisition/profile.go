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
)

func (a *Acquisition) GenerateProfile() error {
	profilePath := filepath.Join(a.StoragePath, "profile.json")
	profile, err := os.Create(profilePath)
	if err != nil {
		return fmt.Errorf("failed to create profile: %v", err)
	}
	defer profile.Close()

	buf, _ := json.MarshalIndent(a, "", "    ")

	profile.WriteString(string(buf[:]))
	profile.Sync()

	return nil
}
