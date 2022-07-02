package api

import (
	"bagogo/modules/util/config"
	"bagogo/modules/util/server"
	"bagogo/routes"
)

func Start(cfg *config.Configuration) error {
	e := server.New()
	routes.SetupRoutes(e, cfg)

	server.Start(e, &server.Config{
		Port: cfg.Server.Port,
		ReadTimeoutSeconds: cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
	})

	return nil
}