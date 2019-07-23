// Copyright (c) 2017-2019 Claudio Guarnieri.
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
	"fmt"
	"github.com/botherder/go-autoruns"
	"github.com/botherder/go-files"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func generateAutoruns() {
	log.Info("Identifying files marked for persistence...")

	// Fetch autoruns.
	autoruns := autoruns.Autoruns()

	// Make backup of autoruns executables.
	for _, autorun := range autoruns {
		if _, err := os.Stat(autorun.ImagePath); err == nil {
			copyName := fmt.Sprintf("%s_%s.bin", autorun.MD5, autorun.ImageName)
			copyPath := filepath.Join(acq.Autoruns, copyName)
			files.Copy(autorun.ImagePath, copyPath)
		}
	}

	// Store the json list to file.
	autorunsJSONPath := filepath.Join(acq.Storage, "autoruns.json")
	autorunsJSON, err := os.Create(autorunsJSONPath)
	if err != nil {
		log.Error("Unable to save autoruns to file: ", err.Error())
		return
	}
	defer autorunsJSON.Close()

	// Encoding into json.
	buf, _ := json.MarshalIndent(autoruns, "", "    ")

	autorunsJSON.WriteString(string(buf[:]))
	autorunsJSON.Sync()

	log.Info("Autoruns collected successfully!")
}
