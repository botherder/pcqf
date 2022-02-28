// pcqf - PC Quick Forensics
// Copyright (c) 2021-2022 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	// "github.com/botherder/pcqf/modules/mac"
	"github.com/botherder/pcqf/modules/crossplatform"
)

func GetModulesList() []Module {
	return []Module{
		crossplatform.GetAutoruns,
		crossplatform.GetProcessList,
		crossplatform.GenerateSystemInfo,
	}
}
