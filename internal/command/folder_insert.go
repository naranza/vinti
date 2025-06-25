// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "fmt"
  "os"
  "path/filepath"
  "vinti/internal/core"
)

func FolderInsert(config *core.Config, folder string) error {
  fullPath := filepath.Join(config.Dir, folder)
  err := os.MkdirAll(fullPath, config.FileModeDir)
  if  err != nil {
    return fmt.Errorf("failed to create directory %q: %w", fullPath, err)
  }
  return nil
}