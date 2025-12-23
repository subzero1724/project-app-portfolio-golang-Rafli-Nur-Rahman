package config

import (
	"os"
	"testing"
)

func TestLoadConfig_EnvOverrides(t *testing.T) {
	os.Setenv("SERVER_PORT", "9999")
	os.Setenv("ENVIRONMENT", "ci")
	defer os.Unsetenv("SERVER_PORT")
	defer os.Unsetenv("ENVIRONMENT")

	cfg := LoadConfig()
	if cfg.ServerPort != "9999" || cfg.Environment != "ci" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestNewDatabase_NoURL(t *testing.T) {
	_, err := NewDatabase(&Config{DatabaseURL: ""})
	if err == nil {
		t.Fatalf("expected error when DATABASE_URL empty")
	}
}
