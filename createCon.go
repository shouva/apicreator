package main

import (
	"log"
	"os"
	"text/template"
)

func createCon() string {
	var strtemplate string
	if config.Database.Server == mySQL {
		strtemplate = `
			package database

			import (
				"fmt"
			
				"github.com/jinzhu/gorm"
			)
			
			// Con : capsuling all
			type Con struct {
				*gorm.DB
			}
			
			// Init : this called when someone call db
			func (db *Con) Init() {
				fmt.Println("init")
			}
			
			// New :
			func New(host, port, username, password, dbname string) (*Con, error) {
				_db, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname)
				if err != nil {
					return &Con{}, err
				}
				err = _db.DB().Ping()
				if err != nil {
					return nil, err
				}
				return &Con{
					DB: _db,
				}, nil
			}	
			`
	} else if config.Database.Server == msSQL {

		strtemplate = `
			package database

			import (
				"fmt"
			
				"github.com/jinzhu/gorm"
			)
			
			// Con : capsuling all
			type Con struct {
				*gorm.DB
			}
			
			// Init : this called when someone call db
			func (db *Con) Init() {
				fmt.Println("init")
			}
			
			// New :
			func New(host, port, username, password, dbname string) (*Con, error) {

				portNumber, _ := strconv.Atoi(port)
				connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
								host, username, password, portNumber, dbname)
				_db, err := gorm.Open("mssql", connectionString)
				if err != nil {
					return &Con{}, err
				}
				err = _db.DB().Ping()
				if err != nil {
					return nil, err
				}
				return &Con{
					DB: _db,
				}, nil
			}	
			`
	}
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := folderdatabase + "/connection.go"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return ""
	}

	// var strout string
	err = tmpl.Execute(f, &API{})

	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	f.Close()
	completer(filename)
	return ""

}
