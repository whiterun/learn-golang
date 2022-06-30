package models

import (
	_ "database/sql"
)

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone *string `json:"phone"`
}

func UserColumns(colname string, usr *User) interface{} {
    switch colname {
	    case "id":
	        return &usr.Id
	    case "name":
	        return &usr.Name
	    case "email":
	        return &usr.Email
        case "phone":
	        return &usr.Phone
	    default:
	        panic("unknown column " + colname)
   }
}