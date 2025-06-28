// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "vinti/internal/core"
)

func FileRename(config *core.Config, dir string, filename string, toDir string) error {
  oldPath := filepath.Join(config.Dir, dir, filename)
  newPath := filepath.Join(config.Dir, toDir, filename)
  return os.Rename(oldPath, newPath)
}
