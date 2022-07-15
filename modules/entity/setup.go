package entity

import (
    "fmt"

    "bagogo/helpers"
    "bagogo/modules/util/config"
	_ "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var logger = helpers.LogInstance()

func Connect(cfg *config.Configuration) {
    dsn := cfg.DB.Username+":"+cfg.DB.Password+"@tcp("+cfg.DB.Host+cfg.DB.Port+")/"+cfg.DB.Database

    database, err := sqlx.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Err", err.Error())
    }

    DB = database
}