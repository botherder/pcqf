package main

import (
	"os"
	"bytes"
	"path/filepath"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(acq)

	profile.WriteString(buf.String())
	profile.Sync()

	log.Info("Profile generated!")
}
