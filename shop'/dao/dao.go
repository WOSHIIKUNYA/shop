package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func ChuShi() {
	m, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		fmt.Println(err)
	}
	database = m
}
