// Package config loads gateway settings from a YAML file, then applies
// environment variable overrides.
package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ListenAddr string `yaml:"listen_addr"` // e.g. ":8080"
	BackendURL string `yaml:"backend_url"` // e.g. "http://localhost:11434"
}

// Load reads the YAML file at path, then overrides fields from the environment:
//
//	GATEWAY_LISTEN_ADDR -> ListenAddr
//	GATEWAY_BACKEND_URL -> BackendURL
//
// TODO:
//  1. `go get gopkg.in/yaml.v3`
//  2. os.ReadFile(path), then yaml.Unmarshal into a Config.
//  3. For each env var above: if os.Getenv returns non-empty, overwrite the field.
//  4. Return an error if BackendURL is still empty (fail at startup, not on first request).
func Load(path string) (Config, error) {
	yml, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	err = yaml.Unmarshal(yml, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	listen_addr := os.Getenv("GATEWAY_LISTEN_ADDR")
	backend_url := os.Getenv("GATEWAY_BACKEND_URL")
	if len(listen_addr) > 0 {
		cfg.ListenAddr = listen_addr
	}
	if len(backend_url) > 0 {
		cfg.BackendURL = backend_url
	}
	if len(cfg.BackendURL) == 0 {
		err = errors.New("backend url is empty")
	}

	return cfg, err
}
