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

func Demo() string {
	/*// Basic Output
	fmt.Println("Hell-O")
	fmt.Printf("%s!\n\n", yolo)

	// Pointer Output
	new := new (string)
	fmt.Println(new)
	fmt.Println(*new)

	// Non-Decimal
	var number = 12345
	fmt .Printf("Non-Decimal: %d \n", number)
	
	// Decimal
	var decimal = 0.1234
	fmt.Printf("Decimal: %.2f \n\n", decimal)

	// Boolean
	var boolean = true
	fmt.Printf("True ? %t \n\n", boolean)

	// Temporary variable
	var point = 8840.0
	if percent := point / 100; percent >= 100 {
	    fmt.Printf("%.1f%s perfect!\n\n", percent, "%")
	} else if percent >= 70 {
	    fmt.Printf("%.1f%s good\n\n", percent, "%")
	} else {
	    fmt.Printf("%.1f%s not bad\n\n", percent, "%")
	}

	// FallThrough
	var lePoint = 6

	switch {
		case lePoint == 8:
		    fmt.Println("perfect")
		case (lePoint < 8) && (lePoint > 3):
		    fmt.Println("awesome")
		    fallthrough
		case lePoint < 5:
		    fmt.Println("you need to learn more")
		default:
	    {
	        fmt.Println("not bad")
	        fmt.Println("you need to learn more")
	    }
	}*/

	return fmt.Sprintf(quote.Go())
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}

	row, err := models.DB.Query("SELECT id, name, email FROM users where id = ?", id)

	if row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			c.AbortWithStatus(404)
		}
		
	    c.JSON(200, gin.H{"data": user})
	} else {
		c.AbortWithStatus(404)
	}

}