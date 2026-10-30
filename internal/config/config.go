// Package config loads gateway settings from a YAML file, then applies
// environment variable overrides.
package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ListenAddr string `yaml:"listen_addr"` // e.g. ":8080"
	BackendURL string `yaml:"backend_url"` // e.g. "http://localhost:11434"
}

func Load(path string) (Config, error) {
	yml, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	err = yaml.Unmarshal(yml, &cfg)
	if err != nil {
		return Config{}, err
	}
	listenAddr := os.Getenv("GATEWAY_LISTEN_ADDR")
	backendURL := os.Getenv("GATEWAY_BACKEND_URL")
	if listenAddr != "" {
		cfg.ListenAddr = listenAddr
	}
	if backendURL != "" {
		cfg.BackendURL = backendURL
	}
	if cfg.BackendURL == "" {
		err = errors.New("backend url is empty")
	}

	return cfg, err
}
