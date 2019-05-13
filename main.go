package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shouva/dailyhelper"
)

var config Config

func main() {
	dailyhelper.ReadConfig(dailyhelper.GetCurrentPath(false)+"/config.json", &config)
	generate()
	defer db.Close()
	// createPostman()
}
