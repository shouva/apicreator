package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Model :
type Model struct {
	*gorm.Model
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
}

func migrateModel() {
	g.AutoMigrate(&Model{})
}

// handler create
func initRoutersModel(r *gin.Engine, urlstring string) {
	r.GET("/"+urlstring+"/", getModels)
	r.GET("/"+urlstring+"/:id", getModel)
	r.POST("/"+urlstring+"", createModel)
	r.PUT("/"+urlstring+"/:id", updateModel)
	r.DELETE("/"+urlstring+"/:id", deleteModel)
}

func getModels(c *gin.Context) {
	var model []Model
	if err := g.Find(&model).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, model)
	}
}

func getModel(c *gin.Context) {
	id := c.Params.ByName("id")
	var model Model
	if err := g.Where("id = ?", id).First(&model).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, model)
	}
}

func createModel(c *gin.Context) {
	var model Model
	c.Bind(&model)
	g.Create(&model)
	c.JSON(200, model)
}

func updateModel(c *gin.Context) {
	var model Model
	id := c.Params.ByName("id")
	if err := g.Where("id = ?", id).First(&model).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.Bind(&model)
	g.Save(&model)
	c.JSON(200, model)
}
func deleteModel(c *gin.Context) {
	id := c.Params.ByName("id")
	var model Model
	d := g.Where("id = ?", id).Delete(&model)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
