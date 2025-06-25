// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "vinti/internal/core"
)

func FileDelete(config *core.Config, dir string, filename string) error {
  fullPath := filepath.Join(config.Dir, dir, filename)

  return  os.Remove(fullPath)
}

