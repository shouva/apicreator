package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	helper "github.com/shouva/dailyhelper"
)

func createConfig() {
	str := `package main

	// Config :
	type Config struct {
		Database Database ` + "`" + `json:"database"` + "`" + `
	}
	
	// Database :
	type Database struct {
		Host     string ` + "`" + `json:"host"` + "`" + `
		Port     string ` + "`" + `json:"port"` + "`" + `
		User     string ` + "`" + `json:"username"` + "`" + `
		Password string ` + "`" + `json:"password"` + "`" + `
		DBName   string ` + "`" + `json:"dbname"` + "`" + `
	}
	`
	filename := helper.GetCurrentPath(false) + "/out/config.go"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	f.Write([]byte(str))
	completer(filename)
}

func copyConfig() {
	sourceFile := helper.GetCurrentPath(false) + "/config.json"
	destinationFile := helper.GetCurrentPath(false) + "/out/config.json"
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}
