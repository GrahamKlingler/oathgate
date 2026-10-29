package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnvOverridesYAML(t *testing.T) {
	path := filepath.Join(t.TempDir(), "gateway.yaml")
	yaml := "listen_addr: \":8080\"\nbackend_url: \"http://from-yaml:11434\"\n"
	if err := os.WriteFile(path, []byte(yaml), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GATEWAY_BACKEND_URL", "http://from-env:11434")

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ListenAddr != ":8080" {
		t.Errorf("ListenAddr = %q, want value from YAML", cfg.ListenAddr)
	}
	if cfg.BackendURL != "http://from-env:11434" {
		t.Errorf("BackendURL = %q, want env override", cfg.BackendURL)
	}
}
