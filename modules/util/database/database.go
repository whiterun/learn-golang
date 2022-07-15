package database

import (
    "bagogo/helpers"
    "bagogo/modules/util/config"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var logger = helpers.LogInstance()

func Connect(cfg *config.Configuration) (*sqlx.DB, error) {
    dsn := cfg.DB.Username+":"+cfg.DB.Password+"@tcp("+cfg.DB.Host+cfg.DB.Port+")/"+cfg.DB.Database

    db, err := sqlx.Open("mysql", dsn)
    if err != nil {
        logger.Println(err.Error())
    }

    return db, nil
}