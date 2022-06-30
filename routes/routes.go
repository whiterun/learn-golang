package routes

import (
	"bagogo/controllers"
	"bagogo/models"
    "bagogo/modules/util/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(cfg *config.Configuration) *gin.Engine {
    r := gin.Default()

    models.Connect(cfg)
    
    api := r.Group("/api")
    api.GET("/user/:id", controllers.GetUser)

    return r
}