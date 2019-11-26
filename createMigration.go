package main

import (
	"log"
	"os"
	"text/template"
)

func createMigrations(models []string) string {
	type params struct {
		Models  []string
		Package string
	}
	strtemplate := `
	package main
	import "{{.Package}}/models"
	func migrate(g *gorm.DB) {
		{{range $index, $model := .Models}}g.AutoMigrate(&models.{{$model}}{})
		{{end}}}
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := folder + "/migrations.go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
		return ""
	}

	// var strout string
	err = tmpl.Execute(f, params{
		Models:  models,
		Package: config.Package,
	})

	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	f.Close()
	completer(filename)
	return ""
}
