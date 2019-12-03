package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	helper "github.com/shouva/dailyhelper"
)

var db *sql.DB

func generate() {
	start := time.Now()

	connect()
	createMain()
	createCon()
	createConfig()
	copyConfig()
	createGitIgnore()
	// handler file harus di lengkapi setelah pembuatan seluruh model. :-)
	createHandlerFile()
	tables := getAllTablename()

	var modelnames []string
	var routes []Route
	c := config.Database
	items := []Item{}

	for _, tablename := range tables {

		modelname := helper.SnakeCaseToCamelCase(helper.Singular(tablename), true)

		tablemodelstring, modelname, queries := createStruct(c.DBName, tablename)
		tablemodelstring = "package models\n" + tablemodelstring
		objectname := helper.SnakeCaseToCamelCase(helper.Singular(tablename), false)
		objectname = helper.BeautySingularity(objectname)
		url := strings.ToLower(objectname)
		objectname = "_" + objectname
		createModel(tablename, tablemodelstring)
		createAPI(tablename, modelname, objectname, url)
		modelnames = append(modelnames, modelname)
		routes = append(routes, Route{Name: modelname, URL: url})
		item := createItems(tablename, url, &queries)
		items = append(items, *item...)

	}
	completer(folderhandler + "/handler.go")

	createMigrations(modelnames)
	createRoutes(routes)
	generatePostman(c.DBName, &items)
	fmt.Println(len(tables), "table * 5 endpoint API have generated!")
	elapsed := time.Since(start)
	log.Printf("Process take finished in %s", elapsed)
}

func createDir() {
	fmt.Println("wow")
	// config.Package
}
func connect() {
	if config.Database.Server == "mysql" {
		connectMYSQL()
	} else if config.Database.Server == "mssql" {
		connectMSSQL()
	} else {
		panic("server tidak dikenali")
	}
	fmt.Println("berhasil connect")
}
func connectMYSQL() {
	c := config.Database
	var err error
	var constring string

	constring = c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DBName + "?&parseTime=True"
	db, err = sql.Open("mysql", constring)

	if err != nil {
		panic(err.Error)
	}
	fmt.Println("berhasil connect")
}
func connectMSSQL() {
	c := config.Database
	var err error
	port, _ := strconv.Atoi(c.Port)
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		c.Host, c.User, c.Password, port, c.DBName)
	fmt.Println(connectionString)
	db, err = sql.Open("mssql", connectionString)
	if err != nil {
		panic(err.Error)
	}
}

const mySQL = "mysql"
const msSQL = "mssql"

func getAllTablename() []string {

	tables := []string{}
	var query string
	if config.Database.Server == msSQL {
		query = `
		SELECT	TABLE_NAME
		  FROM	INFORMATION_SCHEMA.TABLES;
		  `
	} else if config.Database.Server == mySQL {
		query = "SHOW TABLES"
	}

	res, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	var table string

	for res.Next() {
		res.Scan(&table)
		tables = append(tables, table)
	}
	return tables
}

func createFunctionName(strucName, tablename string) string {
	return fmt.Sprintf("\n// TableName : \nfunc (*%s) TableName() string {"+
		"\nreturn \"%s\""+
		"\n}", strucName, tablename)
}

func stringifyType(colType string) string {
	switch colType {
	case "tinyint", "int", "smallint", "mediumint":
		return "uint16"
	case "bigint":
		return "uint64"
	case "char", "enum", "varchar", "nvarchar", "longtext", "mediumtext", "text", "tinytext":
		return "string"
	case "date", "datetime", "time", "timestamp":
		return "*time.Time"
	case "decimal", "double", "numeric":
		return "float64"
	case "float":
		return "float32"
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return "[]byte"
	}
	return ""
}
