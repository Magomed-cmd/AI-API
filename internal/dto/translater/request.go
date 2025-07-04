package translater

type TranslateRequest struct {
	Text         string `json:"text" binding:"required"`
	FromLanguage string `json:"from_language" binding:"required"`
	ToLanguage   string `json:"to_language" binding:"required"`
}
