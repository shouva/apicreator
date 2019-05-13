package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"selamat": "malam"})
	})
}
