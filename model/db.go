package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goDemo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}
