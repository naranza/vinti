// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"gitlab.com/naranza/cogo"
)

type Config struct {
  Dir string
  IncrementMax int
	FileModeDir os.FileMode
	FileModeFile os.FileMode
  IncrementDigits int
	// Add these for SSL
  TlsCertPath string
  TlsKeyPath string
  ServerPort int
	TokenExpiresIn int
}

func DefaultConfig() *Config {
	maxVal := 999999
	return &Config{
		Dir: filepath.Join(os.TempDir(), "vinti"),
		IncrementMax: maxVal,
		IncrementDigits: len(strconv.Itoa(maxVal)),
		FileModeDir: 0700,
		FileModeFile: 0700,
		TlsCertPath: "",
		TlsKeyPath: "",
		ServerPort: 20201,
		TokenExpiresIn 3600,
	}
}

func ConfigLoad(path string) (*Config, error) {
	config := DefaultConfig()
	_, err := os.Stat(path)
	if os.IsExist(err) {
	  err := cogo.LoadConfig(path, &config)
		if err != nil {
			return nil, fmt.Errorf("Failed to load config: %w", err)
		}
	}
	return config, nil
}