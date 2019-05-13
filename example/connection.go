package main

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var g *gorm.DB

func open(host, port, dbname, username, password string) (err error) {
	g, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	return
}
