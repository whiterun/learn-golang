package routes

import (
	"bagogo/controllers"
	"bagogo/models"
    "bagogo/modules/util/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine, cfg *config.Configuration) {
    models.Connect(cfg)
    
    api := e.Group("/api")
    api.GET("/user/:id", controllers.GetUser)
}