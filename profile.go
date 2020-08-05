// Copyright (c) 2017-2020 Claudio Guarnieri.
//
// This file is part of Snoopdigg.
//
// Snoopdigg is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Snoopdigg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Snoopdigg.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func generateProfile() {
	log.Info("Generating profile...")

	// Store the json list to file.
	profilePath := filepath.Join(acq.Storage, "profile.json")
	profile, err := os.Create(profilePath)
	if err != nil {
		log.Error("Unable to create profile: ", err.Error())
		return
	}
	defer profile.Close()

	// Encoding into json.
	buf, _ := json.MarshalIndent(acq, "", "    ")

	profile.WriteString(string(buf[:]))
	profile.Sync()

	log.Info("Profile generated!")
}
