package main

import (
	"log"
	"os"
	"text/template"

	helper "github.com/shouva/dailyhelper"
)

// API struct
type API struct {
	Modelname  string
	Objectname string
	Urlstring  string
	Prefix     string
}

func createAPI(modelname, objectname, urlstring, prefix string) string {
	strtemplate := `
	{{.Prefix}}

	// handler create
	func initRouters{{.Modelname}}(r *gin.Engine, {{.Urlstring}} string) {
		route := r.Group("{{.Urlstring}}")
		route.GET("/", get{{.Modelname}}s)
		route.GET("/:id", get{{.Modelname}})
		route.POST("/", create{{.Modelname}})
		route.PUT("/:id", update{{.Modelname}})
		route.DELETE("/:id", delete{{.Modelname}})
	}
	
	func get{{.Modelname}}s(c *gin.Context) {
		var {{.Objectname}}s []{{.Modelname}}
		if err := g.Find(&{{.Objectname}}s).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, {{.Objectname}}s)
		}
	}
	
	func get{{.Modelname}}(c *gin.Context) {
		id := c.Params.ByName("id")
		var {{.Objectname}} {{.Modelname}}
		if err := g.Where("id = ?", id).First(&{{.Objectname}}).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, {{.Objectname}})
		}
	}
	
	func create{{.Modelname}}(c *gin.Context) {
		var {{.Objectname}} {{.Modelname}}
		c.Bind(&{{.Objectname}})
		g.Create(&{{.Objectname}})
		c.JSON(200, {{.Objectname}})
	}
	
	func update{{.Modelname}}(c *gin.Context) {
		var {{.Objectname}} {{.Modelname}}
		id := c.Params.ByName("id")
		if err := g.Where("id = ?", id).First(&{{.Objectname}}).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.Bind(&{{.Objectname}})
		g.Save(&{{.Objectname}})
		c.JSON(200, {{.Objectname}})
	}
	func delete{{.Modelname}}(c *gin.Context) {
		id := c.Params.ByName("id")
		var {{.Objectname}} {{.Modelname}}
		d := g.Where("id = ?", id).Delete(&{{.Objectname}})
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
	
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}
	api := &API{
		Objectname: objectname,
		Modelname:  modelname,
		Urlstring:  urlstring,
		Prefix:     prefix,
	}

	// openfile
	filename := helper.GetCurrentPath(false) + "/out/" + modelname + ".go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
		return ""
	}

	// var strout string
	err = tmpl.Execute(f, api)

	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	f.Close()
	completer(filename)
	return ""
}
