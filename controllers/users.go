package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rsc.io/quote"
	"bagogo/helpers"
	"bagogo/models"
)

var logger = helpers.LogInstance()
var yolo = "Lorem Ipsum"

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	var columns []string

	row, _ := models.DB.Query("SELECT id, name, email, phone FROM users where id = ?", id)

	if row.Next() {
		columns, _ = row.Columns()
		colNum := len(columns)

		cols := make([]interface{}, colNum)
        for i := 0; i < colNum; i++ {
        	_, found := helpers.FindSlice(columns, columns[i])

        	if found {
	            cols[i] = models.UserColumns(columns[i], &user)
	        }
        }

		if err := row.Scan(cols...); err != nil {
			c.AbortWithStatus(404)
		}
		
	    c.JSON(200, gin.H{"data": user})
	} else {
		c.AbortWithStatus(404)
	}
}

func Demo() string {
	return fmt.Sprintf(quote.Go())
}