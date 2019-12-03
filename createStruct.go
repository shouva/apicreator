package main

import (
	"database/sql"
	"fmt"

	"github.com/icrowley/fake"
	helper "github.com/shouva/dailyhelper"
)

func createStruct(database, table string) (string, string, []Query) {
	var query string
	if config.Database.Server == mySQL {
		query = `SELECT 	COLUMN_NAME, 
						DATA_TYPE, 
						IS_NULLABLE, 
						COLUMN_KEY
			FROM 		INFORMATION_SCHEMA.COLUMNS 
			WHERE 		TABLE_SCHEMA = ? 
			AND 		TABLE_NAME = ?`
	} else if config.Database.Server == msSQL {
		query = `		
		SELECT c.COLUMN_NAME, c.DATA_TYPE, c.IS_NULLABLE, COALESCE(p.CONSTRAINT_TYPE,'') from INFORMATION_SCHEMA.COLUMNS c LEFT JOIN (
			SELECT  ccu.TABLE_CATALOG,tc.TABLE_NAME, ccu.COLUMN_NAME, tc.CONSTRAINT_TYPE from INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc,
			INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE ccu
			WHERE tc.CONSTRAINT_NAME = ccu.CONSTRAINT_NAME 
			AND tc.CONSTRAINT_CATALOG=? 
			AND ccu.TABLE_NAME=?
		) as p on p.TABLE_NAME = c.TABLE_NAME
		AND p.COLUMN_NAME = c.COLUMN_NAME
		WHERE c.TABLE_CATALOG = ?
		AND c.TABLE_NAME = ?
		`
	}
	var rows *sql.Rows
	var err error
	if config.Database.Server == mySQL {
		rows, err = db.Query(query, database, table)
	} else if config.Database.Server == msSQL {
		rows, err = db.Query(query, database, table, database, table)
	}

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return "", "", nil
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return "", "", nil
	}
	queries := []Query{}
	structName := helper.SnakeCaseToCamelCase(helper.Singular(table), true)
	structName = helper.BeautySingularity(structName)
	strStuct := "package model "
	strStuct = "\n// " + structName + " : "
	strStuct += "\ntype " + structName + " struct {\n"
	for rows.Next() {
		var colname string
		var colKey string
		var colType string
		var nullable string
		err = rows.Scan(&colname, &colType, &nullable, &colKey)
		if err != nil {
			fmt.Println(err)
		}
		field := helper.SnakeCaseToCamelCase(colname, true)
		field = helper.BeautySingularity(field)
		colType = stringifyType(colType)
		if colname == "id" {
			field = "ID"
		} else {
			if colType == "string" {
				queries = append(queries, Query{
					Key:   colname,
					Value: fake.FirstName(),
				})
			} else {
				queries = append(queries, Query{
					Key:   colname,
					Value: fake.DigitsN(2),
				})
			}
		}
		if colKey == "PRI" || colKey == "PRIMARY KEY" {
			strStuct += fmt.Sprintf("\n\t %s %s `gorm:\"column:%s;primary_key\" form:\"%s;primary_key\" json:\"%s,omitempty;primary_key\" bson:\"%s\" query:\"%s\" form:\"%s\" xml:\"%s\"`", field, colType, colname, colname, colname, colname, colname, colname, colname)
		} else {
			strStuct += fmt.Sprintf("\n\t %s %s `gorm:\"column:%s\" form:\"%s\" json:\"%s,omitempty\" bson:\"%s\" query:\"%s\" form:\"%s\" xml:\"%s\"`", field, colType, colname, colname, colname, colname, colname, colname, colname)

		}
	}
	strStuct += "\n}"
	strStuct += createFunctionName(structName, table)

	return strStuct, structName, queries
}
