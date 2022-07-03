package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tienanh129902/go-rest-api/config"
	"github.com/tienanh129902/go-rest-api/handlers"
	"github.com/tienanh129902/go-rest-api/routers"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

// @title Quiz game API
// @version 1.0
// @description This is a basic API for a quiz game using Gin and Gorm.
// @schemes http https
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := handlers.InitializeHandler(); err != nil {
		log.Fatalln(err)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        routers.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
