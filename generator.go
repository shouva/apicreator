package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/icrowley/fake"

	helper "github.com/shouva/dailyhelper"
)

var db *sql.DB

func generate() {
	connect()
	tables := getAllTablename()
	createMain()
	createCon()
	var modelnames []string
	var routes []Route
	c := config.Database
	items := []Item{}
	for _, tablename := range tables {
		modelname := helper.SnakeCaseToCamelCase(helper.Singular(tablename), true)
		prefix, queries := createStruct(c.DBName, tablename)
		prefix = "package main\n" + prefix
		// fmt.Println(prefix)
		objectname := helper.SnakeCaseToCamelCase(helper.Singular(tablename), false)
		url := strings.ToLower(objectname)

		createAPI(modelname, objectname, url, prefix)
		modelnames = append(modelnames, modelname)
		routes = append(routes, Route{Name: modelname, URL: url})
		item := createItems(tablename, url, &queries)
		items = append(items, *item...)
	}
	createMigrations(modelnames)
	createRoutes(routes)
	generatePostman(c.DBName, &items)
}
func connect() {
	c := config.Database
	var err error
	constring := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DBName + "?&parseTime=True"
	db, err = sql.Open("mysql", constring)
	if err != nil {
		panic(err.Error)
	}
}
func getAllTablename() []string {

	tables := []string{}
	res, err := db.Query("SHOW TABLES")

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

func createStruct(database, table string) (string, []Query) {
	query := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"
	rows, err := db.Query(query, database, table)
	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return "", nil
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return "", nil
	}
	queries := []Query{}
	structName := helper.SnakeCaseToCamelCase(helper.Singular(table), true)
	strStuct := "\n// " + structName + " : "
	strStuct += "\ntype " + structName + " struct {\n"
	for rows.Next() {
		var colname string
		var colKey string
		var colType string
		var nullable string
		rows.Scan(&colname, &colKey, &colType, &nullable)
		field := helper.SnakeCaseToCamelCase(colname, true)
		if colname == "id" {
			field = "ID"
		} else {
			queries = append(queries, Query{
				Key:   colname,
				Value: fake.DigitsN(10),
			})
		}
		colType = stringifyType(colType)
		if colKey == "PRI" {
			strStuct += fmt.Sprintf("\n\t %s %s `gorm:\"column:%s;primary_key\" form:\"%s;primary_key\" json:\"%s;primary_key\"`", field, colType, colname, colname, colname)
		} else {
			strStuct += fmt.Sprintf("\n\t %s %s `gorm:\"column:%s\" form:\"%s\" json:\"%s\"`", field, colType, colname, colname, colname)

		}
	}
	strStuct += "\n}"
	strStuct += createFunctionName(structName, table)
	return strStuct, queries
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
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		return "string"
	case "date", "datetime", "time", "timestamp":
		return "*time.Time"
	case "decimal", "double":
		return "float64"
	case "float":
		return "float32"
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return "[]byte"
	}
	return ""
}
