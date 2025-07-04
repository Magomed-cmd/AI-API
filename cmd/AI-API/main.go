package main

import (
	"AI-API/internal/config"
	"AI-API/internal/routes"
	"AI-API/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Передаем конфиг в сервис
	translatorService := service.NewTranslatorService(cfg)

	routes.RegisterRoutes(r, translatorService)

	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
