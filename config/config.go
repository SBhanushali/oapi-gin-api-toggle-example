package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FeatureFlags map[string]bool `yaml:"feature_flags"`
}

func LoadConfig() (*Config, error) {
	var config Config

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(dir, "config", "config.yaml")
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
