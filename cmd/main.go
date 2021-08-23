// pcqf - PC Quick Forensics
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package main

import (
	"fmt"
	"os"

	"github.com/botherder/pcqf/acquisition"
	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/mem"
	"github.com/i582/cfmt/cmd/cfmt"
)

func printError(desc string, err error) {
	cfmt.Printf("{{ERROR:}}::red|bold %s: {{%s}}::italic\n",
		desc, err.Error())
}

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

func main() {
	acq, err := acquisition.New()
	if err != nil {
		cfmt.Println(err)
		return
	}

	cfmt.Printf("Started acquisition {{%s}}::magenta|underline\n", acq.UUID)

	fmt.Println("Generating system profile...")
	err = acq.GenerateProfile()
	if err != nil {
		printError("Failed to generate system profile", err)
	}

	fmt.Println("Generating process list...")
	err = acq.GenerateProcessList()
	if err != nil {
		printError("Failed to generate process list", err)
	}

	fmt.Println("Generating list of persistent software...")
	err = acq.GenerateAutoruns()
	if err != nil {
		printError("Failed to generate list of persistent software", err)
	}

	virt, _ := mem.VirtualMemory()
	virtTotal := virt.Total / (1000 * 1000 * 1000)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Take a memory snapshot (it would take circa %d GB of space)", virtTotal),
		IsConfirm: true,
	}

	takeMemory, err := prompt.Run()
	if err == nil && takeMemory == "y" {
		acq.GenerateMemoryDump()
	} else {
		fmt.Println("Skipping memory acquisition.")
	}

	err = acq.StoreSecurely()
	if err != nil {
		printError("Something failed while encrypting the acquisition", err)
		cfmt.Println("{{WARNING: The secure storage of the acquisition folder failed! The data is unencrypted!}}::red|bold")
	}

	fmt.Println("Acquisition completed.")

	fmt.Println("Press Enter to finish ...")
	var b = make([]byte, 1)
	os.Stdin.Read(b)
}
