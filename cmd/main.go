// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"fmt"
	"os"

	"github.com/botherder/pcqf/acquisition"
	"github.com/i582/cfmt/cmd/cfmt"
)

var acq *acquisition.Acquisition

func init() {
	cfmt.Print(`
	    {{                        ____  }}::green
	    {{      ____  _________  / __/  }}::yellow
	    {{     / __ \/ ___/ __ '/ /_    }}::red
	    {{    / /_/ / /__/ /_/ / __/    }}::magenta
	    {{   / .___/\___/\__, /_/       }}::blue
	    {{  /_/            /_/          }}::cyan
	`)
	cfmt.Println("\t\tpcqf - PC Quick Forensics")
	cfmt.Println()
}

func printError(msg string, err error) {
	cfmt.Printf("{{ERROR:}}::red|bold %s: {{%s}}::italic\n",
		msg, err.Error())
}

func systemPause() {
	cfmt.Println("Press {{Enter}}::bold|green to finish ...")
	os.Stdin.Read(make([]byte, 1))
}

func main() {
	var err error
	acq, err = acquisition.New()
	if err != nil {
		cfmt.Println(err)
		return
	}

	cfmt.Printf("Started acquisition {{%s}}::magenta|underline\n", acq.UUID)

	RunModules()

	err = acq.StoreSecurely()
	if err != nil {
		printError("Something failed while encrypting the acquisition", err)
		cfmt.Println("{{WARNING: The secure storage of the acquisition folder failed! The data is unencrypted!}}::red|bold")
	}

	fmt.Println("Acquisition completed.")

	systemPause()
}
