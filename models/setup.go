package models

import (
    "fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ali_cm")
    if err != nil {
        fmt.Println("Err", err.Error())
    }

    DB = database
}