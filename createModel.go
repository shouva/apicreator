package main

import (
	"log"
	"os"
)

func createModel(tablename, _struct string) {
	filename := foldermodel + "/" + tablename + ".go"

	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
	}
	f.WriteString(_struct)
	f.Close()
	completer(filename)
}
