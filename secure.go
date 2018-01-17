package main

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/botherder/go-files"
	log "github.com/Sirupsen/logrus"
)

func storeSecurely() {
	publicKeyFile := filepath.Join(getCwd(), "public.asc")
	if _, err := os.Stat(publicKeyFile); os.IsNotExist(err) {
		return
	}

	log.Info("You provided a PGP public key, storing the acquisition securely.")

	zipFileName := fmt.Sprintf("%s.zip", acq.UUID)
	zipFilePath := filepath.Join(getCwd(), "acquisitions", zipFileName)

	log.Info("Compressing the acquisition folder. This might take a while...")

	err := files.Zip(acq.Storage, zipFilePath)
	if err != nil {
		log.Error("Something failed compressing the acquisition: ", err.Error())
		log.Warning("The secure storage of the acquisition folder failed! The data is unencrypted!")
		return
	}

	log.Info("Encrypting the compressed archive. This might take a while...")

	// TODO: encrypt file.
	// TODO: delete folder.
}
