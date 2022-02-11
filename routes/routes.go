package routes

import "bagogo/controllers"

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": controllers.Demo(),
		})
	})

    return r
}