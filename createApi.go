package main

import (
	"log"
	"os"
	"text/template"
)

// API struct
type API struct {
	Modelname  string
	Objectname string
	Urlstring  string
	Prefix     string
}

func createAPI(tablename, modelname, objectname, urlstring string) string {
	strtemplate := `
	package handlers

	// InitRouters{{.Modelname}} : 
	func (h *Handler) InitRouters{{.Modelname}}(r *gin.Engine, {{.Objectname}} string) {
		route := r.Group("{{.Urlstring}}")
		route.GET("/", h.get{{.Modelname}}s)
		route.GET("/:id", h.get{{.Modelname}})
		route.POST("/", h.create{{.Modelname}})
		route.PUT("/:id", h.update{{.Modelname}})
		route.DELETE("/:id", h.delete{{.Modelname}})
	}
	
	func (h *Handler) get{{.Modelname}}s(c *gin.Context) {
		var {{.Objectname}}s []models.{{.Modelname}}
		if err := h.db.Find(&{{.Objectname}}s).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, {{.Objectname}}s)
		}
	}
	
	func (h *Handler) get{{.Modelname}}(c *gin.Context) {
		id := c.Params.ByName("id")
		var {{.Objectname}} models.{{.Modelname}}
		if err := h.db.Where("id = ?", id).First(&{{.Objectname}}).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, {{.Objectname}})
		}
	}
	
	func (h *Handler) create{{.Modelname}}(c *gin.Context) {
		var {{.Objectname}} models.{{.Modelname}}
		c.BindJSON(&{{.Objectname}})
		h.db.Create(&{{.Objectname}})
		c.JSON(200, {{.Objectname}})
	}
	
	func (h *Handler) update{{.Modelname}}(c *gin.Context) {
		var {{.Objectname}} models.{{.Modelname}}
		id := c.Params.ByName("id")
		if err := h.db.Where("id = ?", id).First(&{{.Objectname}}).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.BindJSON(&{{.Objectname}})
		h.db.Save(&{{.Objectname}})
		c.JSON(200, {{.Objectname}})
	}
	func (h *Handler) delete{{.Modelname}}(c *gin.Context) {
		id := c.Params.ByName("id")
		var {{.Objectname}} models.{{.Modelname}}
		d := h.db.Where("id = ?", id).Delete(&{{.Objectname}})
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
	}

	// openfile
	filename := folderhandler + "/" + urlstring + ".go"
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
