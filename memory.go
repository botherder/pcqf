package main

import (
	"os"
	"path/filepath"
)

var binPath string = filepath.Join(getCwd(), "bin")

func initBinFolder() error {
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		err = os.MkdirAll(binPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
