package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_web_native")
	if err != nil {
	panic(err)

	log.Println("Database connected")
	DB = db
}
}