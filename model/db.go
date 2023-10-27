package model

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func Setup() {
	dsn := flag.String("dsn", os.Getenv("DSN"), "root:@tcp(127.0.0.1:3306)/goDemo")
	flag.Parse()

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	fmt.Println("Success!")
}
