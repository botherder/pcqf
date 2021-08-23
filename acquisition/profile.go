// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func (a *Acquisition) GenerateProfile() {
	log.Info("Generating profile...")

	profilePath := filepath.Join(a.StoragePath, "profile.json")
	profile, err := os.Create(profilePath)
	if err != nil {
		log.Error("Unable to create profile: ", err.Error())
		return
	}
	defer profile.Close()

	buf, _ := json.MarshalIndent(a, "", "    ")

	profile.WriteString(string(buf[:]))
	profile.Sync()

	log.Info("Profile generated!")
}
