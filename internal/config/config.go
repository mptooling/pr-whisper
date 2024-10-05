package config

import (
	"github.com/mptooling/pr-whisper/internal/domain"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	configPath string
}

func NewConfig(configPath string) *Config {
	return &Config{
		configPath: configPath,
	}
}

func (c Config) loadConfig() (*domain.WhisperConfig, error) {
	data, err := os.ReadFile(c.configPath)
	if err != nil {
		return nil, err
	}

	var config domain.WhisperConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
