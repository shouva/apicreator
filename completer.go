package main

import (
	"log"
	"os"
	"os/exec"
)

func completer(filename string) {
	out, err := exec.Command("goimports", filename).Output()
	// out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filename)
	f.Write(out)
	f.Close()
}
