package api

import (
	"log"
	"bagogo/modules/util/config"
	"bagogo/routes"
)

func Start(cfg *config.Configuration) error {
	r := routes.SetupRoutes()

    log.Fatal(r.Run("localhost:3000"))
	
	return nil
}