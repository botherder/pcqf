package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/botherder/go-autoruns"
	"github.com/botherder/go-files"
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
