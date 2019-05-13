package main

import (
	"log"
	"os"

	helper "github.com/shouva/dailyhelper"
)

func createGitIgnore() {
	str := `
	*
	!*.*
	!/**/
	*.exe
	config.json
	`
	filename := helper.GetCurrentPath(false) + "/out/config.go"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	f.Write([]byte(str))
}
