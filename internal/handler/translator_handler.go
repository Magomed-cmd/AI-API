package handler

import (
	"AI-API/internal/dto/translater"
	"AI-API/internal/service"
	"github.com/gin-gonic/gin"
)

type TranslaterHandler struct {
	service *service.TranslatorService
}

func NewTranslaterHandler(s *service.TranslatorService) *TranslaterHandler {
	return &TranslaterHandler{
		service: s,
	}
}

func (t *TranslaterHandler) Translate(c *gin.Context) {
	ctx := c.Request.Context()

	info := translater.TranslateRequest{}

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	result, err := t.service.Translate(ctx, info.Text, info.FromLanguage, info.ToLanguage)
	if err != nil {
		c.JSON(500, gin.H{"error": "Translation failed", "details": err.Error()})
		return
	}

	c.JSON(
		200, gin.H{
			"result": result,
		})

}

func (t *TranslaterHandler) GetLanguages(c *gin.Context) {
	languages := t.service.GetSupportedLanguages()

	c.JSON(200, gin.H{
		"languages":           languages,
		"supported_languages": "List of supported languages",
	})
}
