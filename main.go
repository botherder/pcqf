package main

import (
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
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
	fmt.Println("    (c) 2017-2018 Claudio Guarnieri (nex@nex.sx)      ")
	fmt.Println("                                                      ")

	// Set up the logging.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	acq.Initialize()

	log.Info("Started acquisition ", acq.Folder)

	generateProfile()
	generateProcessList()
	generateAutoruns()
	generateMemoryDump()

	log.Info("Acquisition completed.")

	log.Info("Press Enter to finish ...")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
}
