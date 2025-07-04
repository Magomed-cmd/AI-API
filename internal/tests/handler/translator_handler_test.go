package handler_test

import (
	"AI-API/internal/config"
	"AI-API/internal/handler"
	"AI-API/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetLanguages(t *testing.T) {

	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		Languages: []config.Language{
			{Code: "en", Name: "English"},
			{Code: "ru", Name: "Russian"},
			{Code: "de", Name: "German"},
		},
	}

	translatorService := service.NewTranslatorService(cfg)
	translatorHandler := handler.NewTranslaterHandler(translatorService)

	router := gin.New()
	router.GET("/languages", translatorHandler.GetLanguages)

	req, err := http.NewRequest("GET", "/languages", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Contains(t, response, "languages")

	languages, ok := response["languages"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, languages, 3)

	languageCodes := make([]string, len(languages))
	for i, lang := range languages {
		code, ok := lang.(string)
		assert.True(t, ok)
		languageCodes[i] = code
	}

	assert.Contains(t, languageCodes, "en")
	assert.Contains(t, languageCodes, "ru")
	assert.Contains(t, languageCodes, "de")
}

func TestGetLanguages_EmptyConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		Languages: []config.Language{},
	}

	translatorService := service.NewTranslatorService(cfg)
	translatorHandler := handler.NewTranslaterHandler(translatorService)

	router := gin.New()
	router.GET("/languages", translatorHandler.GetLanguages)

	req, err := http.NewRequest("GET", "/languages", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	languages := response["languages"].([]interface{})
	assert.Len(t, languages, 0)
}
