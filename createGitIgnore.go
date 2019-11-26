package main

import (
	"log"
	"os"
)

func createGitIgnore() {
	str := `*
!*.*
!/**/
*.exe
config.json
	`
	filename := folder + "/.gitignore"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	f.Write([]byte(str))
}
