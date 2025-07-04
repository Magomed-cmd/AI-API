package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port         int           `mapstructure:"port"`
		Host         string        `mapstructure:"host"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
		IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	} `mapstructure:"server"`

	OpenRouter struct {
		APIKey     string        `mapstructure:"api_key"`
		BaseURL    string        `mapstructure:"base_url"`
		Model      string        `mapstructure:"model"`
		Timeout    time.Duration `mapstructure:"timeout"`
		MaxRetries int           `mapstructure:"max_retries"`
		RetryDelay time.Duration `mapstructure:"retry_delay"`
	} `mapstructure:"openrouter"`

	Translation struct {
		MaxTextLength      int     `mapstructure:"max_text_length"`
		DefaultTemperature float64 `mapstructure:"default_temperature"`
		MaxTokens          int     `mapstructure:"max_tokens"`
	} `mapstructure:"translation"`

	Languages []Language `mapstructure:"languages"`
}

type Language struct {
	Code string `mapstructure:"code"`
	Name string `mapstructure:"name"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("TRANSLATOR")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
