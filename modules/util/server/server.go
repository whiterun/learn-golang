package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// New instantates new Gin server
func New() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())

	e.GET("/", healthCheck)

	return e
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Everything is Ok :)")
}

// Config represents server specific config
type Config struct {
	Port string
	ReadTimeoutSeconds int
	WriteTimeoutSeconds int
	Debug bool
}

// Start starts gin server
func Start(e *gin.Engine, cfg *Config) {
	srv := &http.Server{
		Addr: cfg.Port,
		Handler: e,
		ReadTimeout: time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
