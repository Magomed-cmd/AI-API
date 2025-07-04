package service

import (
	"AI-API/internal/config"
	"AI-API/internal/dto/translater"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TranslatorService struct {
	cfg *config.Config
}

func NewTranslatorService(cfg *config.Config) *TranslatorService {
	return &TranslatorService{cfg: cfg}
}

func (t *TranslatorService) Translate(ctx context.Context, text, fromLanguage, toLanguage string) (string, error) {

	url := t.cfg.OpenRouter.BaseURL + "/chat/completions"

	content := fmt.Sprintf(`Translate this text from %s to %s: "%s"
		
		Respond with ONLY valid JSON in this format:
		{"translated_text": "your translation here"}
		
		Do not add any explanations, markdown, or extra text.`, fromLanguage, toLanguage, text)

	payload := strings.NewReader(fmt.Sprintf(`{
		  "model": "%s",
		  "messages": [
			  {
				 "role": "user", 
				 "content": %q
			  }
		  ]
		}`, t.cfg.OpenRouter.Model, content))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.cfg.OpenRouter.APIKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}

	defer func() { _ = res.Body.Close() }()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var response translater.OpenRouterResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from AI")
	}

	var translationResult struct {
		TranslatedText string `json:"translated_text"`
	}

	if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &translationResult); err != nil {
		return "", fmt.Errorf("failed to parse translation: %w", err)
	}

	return translationResult.TranslatedText, nil
}

func (t *TranslatorService) GetSupportedLanguages() []string {
	codes := make([]string, len(t.cfg.Languages))
	for i, lang := range t.cfg.Languages {
		codes[i] = lang.Code
	}
	return codes
}
