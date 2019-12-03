package main

// Config :
type Config struct {
	Database Database `json:"database"`
	Package  string   `json:"package"`
}

// Database :
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Server   string `json:"server"`
}
