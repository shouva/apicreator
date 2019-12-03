package main

import (
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/shouva/dailyhelper"
)

var config Config
var folder string
var foldermodel string
var folderhandler string
var folderdatabase string

func main() {
	dailyhelper.ReadConfig(dailyhelper.GetCurrentPath(false)+"/config.json", &config)
	folder = os.Getenv("GOPATH") + "/src/" + config.Package
	fmt.Println(folder)
	foldermodel = folder + "/models"
	folderhandler = folder + "/handlers"
	folderdatabase = folder + "/database"
	os.MkdirAll(folder, os.ModePerm)
	os.MkdirAll(foldermodel, os.ModePerm)
	os.MkdirAll(folderhandler, os.ModePerm)
	os.MkdirAll(folderdatabase, os.ModePerm)
	generate()
	defer db.Close()
	// createPostman()
}
