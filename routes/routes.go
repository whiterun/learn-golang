package routes

import (
	// "bagogo/controllers"
    "bagogo/modules/util/config"
	"bagogo/modules/util/database"

    "bagogo/modules/api/user"
    ut "bagogo/modules/api/user/transport"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine, cfg *config.Configuration) error {
    db, err := database.Connect(cfg)

    if err != nil {
        return err
    }
    
    api := e.Group("/api")
    ut.NewHTTP(user.Initialize(db), api)

    return nil
}