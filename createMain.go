package main

import (
	"log"
	"os"
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
		var config Config
		dailyhelper.ReadConfig(dailyhelper.GetCurrentPath(false)+"/config.json", &config)
		c_db := config.Database
		g, _ = open(c_db.Host, c_db.Port, c_db.DBName, c_db.User, c_db.Password)
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
	completer(filename)
	return ""
}
