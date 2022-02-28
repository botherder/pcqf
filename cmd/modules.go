// pcqf - PC Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"fmt"

	"github.com/botherder/pcqf/acquisition"
	"github.com/botherder/pcqf/utils"
)

type Module func(acq *acquisition.Acquisition) error

func RunModules() {
	for _, mod := range GetModulesList() {
		err := mod(acq)
		if err != nil {
			printError(fmt.Sprintf("Unable to run module %s:",
				utils.GetFunctionName(mod)), err)
		}
	}
}
