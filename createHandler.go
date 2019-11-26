package main

import (
	"log"
	"os"
	"text/template"
)

func createHandlerFile() {
	filecontaint := `
	package handlers

		// Handler :
		type Handler struct {
			db   *database.Con
		}

		// NewHandler :
		func NewHandler(database *database.Con) *Handler {
			handler := &Handler{
				db: database,
			}
			handler.init()
			return handler
		}
		func (h *Handler) init(){
			
		}
		`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(filecontaint)
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	// openfile
	filename := folderhandler + "/handler.go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
	}

	// var strout string
	err = tmpl.Execute(f, config)

	if err != nil {
		log.Fatal("Execute: ", err)
	}
	f.Close()
}
