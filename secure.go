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
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/botherder/go-files"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func logEncryptFail(err error) {
	log.Error("Something failed while encrypting the compressed acquisition: ", err.Error())
	log.Warning("The secure storage of the acquisition folder failed! The data is unencrypted!")
}

func storeSecurely() {
	keyFilePath := filepath.Join(getCwd(), "public.asc")
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
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

	// Prepare for error-handling hell...

	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		logEncryptFail(err)
		return
	}
	defer zipFile.Close()

	keyFile, _ := os.Open(keyFilePath)
	if err != nil {
		logEncryptFail(err)
		return
	}
	defer keyFile.Close()

	block, err := armor.Decode(keyFile)
	if err != nil {
		logEncryptFail(err)
		return
	}

	key, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		logEncryptFail(err)
		return
	}

	encFileName := fmt.Sprintf("%s.enc", zipFileName)
	encFilePath := filepath.Join(getCwd(), "acquisitions", encFileName)

	dst, err := os.Create(encFilePath)
	if err != nil {
		logEncryptFail(err)
		return
	}
	defer dst.Close()

	cryptor, err := openpgp.Encrypt(dst, []*openpgp.Entity{key}, nil, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		logEncryptFail(err)
		return
	}

	_, err = io.Copy(cryptor, zipFile)
	if err != nil {
		logEncryptFail(err)
		return
	}

	// We need to explicitly close this before being able to delete it.
	zipFile.Close()
	cryptor.Close()

	// Unbelievable, we managed to get here.

	log.Info("Acquisition successfully encrypted at ", encFilePath)

	// TODO: we should securely wipe the files.

	err = os.Remove(zipFilePath)
	if err != nil {
		log.Error("Failed to delete the unencrypted compressed archive: ", err.Error())
		log.Warning("The secure storage of the acquisition folder failed! The data is unencrypted!")
	}

	err = os.RemoveAll(acq.Storage)
	if err != nil {
		log.Error("Failed to delete the original unencrypted acquisition folder: ", err.Error())
		log.Warning("The secure storage of the acquisition folder failed! The data is unencrypted!")
	}
}
