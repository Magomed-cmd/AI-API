package translater

type TranslateResponse struct {
	TranslatedText string `json:"translated_text"`
	FromLanguage   string `json:"from_language"`
	ToLanguage     string `json:"to_language"`
	Success        bool   `json:"success"`
}

var response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
