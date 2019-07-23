// Copyright (c) 2017-2019 Claudio Guarnieri.
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
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"github.com/shirou/gopsutil/mem"
	"os"
	"strings"
)

var acq Acquisition

func main() {
	fmt.Println("                                 _ _                  ")
	fmt.Println("                                | (_)                 ")
	fmt.Println(" ___ _ __   ___   ___  _ __   __| |_  __ _  __ _      ")
	fmt.Println("/ __| '_ \\ / _ \\ / _ \\| '_ \\ / _` | |/ _` |/ _` | ")
	fmt.Println("\\__ \\ | | | (_) | (_) | |_) | (_| | | (_| | (_| |   ")
	fmt.Println("|___/_| |_|\\___/ \\___/| .__/ \\__,_|_|\\__, |\\__, |")
	fmt.Println("                      | |             __/ | __/ |     ")
	fmt.Println("                      |_|            |___/ |___/      ")
	fmt.Println("    (c) 2017-2019 Claudio Guarnieri (nex@nex.sx)      ")
	fmt.Println("                                                      ")

	// Set up the logging.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	acq.Initialize()

	log.Info("Started acquisition ", acq.Folder)

	generateProfile()
	generateProcessList()
	generateAutoruns()

	virt, _ := mem.VirtualMemory()
	total := virt.Total / (1000 * 1000 * 1000)

	log.Warning("Do you want to take a memory snapshot (it will take circa ", total, " GB of space) ? [y/N]")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToLower(choice))
	if choice == "y" {
		generateMemoryDump()
	} else {
		log.Info("Skipping memory acquisition.")
	}

	storeSecurely()

	log.Info("Acquisition completed.")
	log.Info("Press Enter to finish ...")
	var b = make([]byte, 1)
	os.Stdin.Read(b)
}
