// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"os"
	"path/filepath"

	"github.com/botherder/pcqf/utils"
)

var binPath = filepath.Join(utils.GetCwd(), "bin")

func initBinFolder() error {
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		err = os.MkdirAll(binPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
