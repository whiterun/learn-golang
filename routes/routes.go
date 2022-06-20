package routes

import (
	"bagogo/controllers"
	"github.com/gin-gonic/gin"
	"bagogo/models"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    models.Connect()
    
    r.GET("/user/:id", controllers.GetUser)

    return r
}