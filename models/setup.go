package models

import (
    "fmt"

    "bagogo/helpers"
    "bagogo/modules/util/config"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var logger = helpers.LogInstance()

func Connect(cfg *config.Configuration) {
    dsn := cfg.DB.Username+":"+cfg.DB.Password+"@tcp("+cfg.DB.Host+cfg.DB.Port+")/"+cfg.DB.Database

    database, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err", err.Error())
    }

    DB = database
}