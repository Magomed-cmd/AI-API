package routes

import (
	"AI-API/internal/handler"
	"AI-API/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *service.TranslatorService) {

	translater := handler.NewTranslaterHandler(s)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the AI-API",
		})
	})

	r.POST("/translate", translater.Translate)
	r.GET("/languages", translater.GetLanguages)
}