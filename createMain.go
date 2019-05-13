package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"

	helper "github.com/shouva/dailyhelper"
)

func createMain() string {
	strtemplate := `
	package main

	import (
		"github.com/gin-gonic/gin"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"

	)
	func main() {
		// connect to db
		g, _ = open("localhost", "3306", "testingsaja", "username", "password")
		defer g.Close()
		migrate(g)

		r := gin.New()
		gin.SetMode(gin.DebugMode)
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"selamat": "malam"})
		})
		loadrouter(r)
		r.Run()
	}
	
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := helper.GetCurrentPath(false) + "/out/main.go"
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
