package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"

	helper "github.com/shouva/dailyhelper"
)

func createCon() string {
	strtemplate := `
	package main

	import "github.com/jinzhu/gorm"
	
	var g *gorm.DB
	
	func open(host, port, dbname, username, password string) (oGorm *gorm.DB, err error) {
		oGorm, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
		return
	}
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := helper.GetCurrentPath(false) + "/out/con.go"
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
	out, err := exec.Command("goimports", filename).Output()
	// out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Create(filename)
	f.Write(out)
	f.Close()
	return ""

}
