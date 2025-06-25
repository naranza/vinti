// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "vinti/internal/core"
)

func FileRead(config *core.Config, dir string, filename string) (string, error) {
  fullPath := filepath.Join(config.Dir, dir, filename)

  data, err := os.ReadFile(fullPath)
  if err != nil {
    return "", err
  }

  return string(data), nil
}
