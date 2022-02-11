package controllers

import "fmt"

import models "bagogo/models"

func SqlQuery() {
    db, err := models.Connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    rows, err := db.Query("select id, name, email from users limit 2")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []models.User

    for rows.Next() {
        var each = models.User{}
        var err = rows.Scan(&each.Id, &each.Name, &each.Email)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
        fmt.Println(each.Name)
    }
}
