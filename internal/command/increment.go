// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "fmt"
  "os"
  "path/filepath"
	"vinti/internal/core"
)

func IncrementFile(config *core.Config, dir, name string) (string, *os.File, error) {
	for i := 0; i <= config.IncrementMax; i++ {
		filename := fmt.Sprintf("%s%0*d", name, config.IncrementDigits, i)
		path := filepath.Join(config.Dir, dir, filename)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_SYNC, config.FileModeFile)
			if err == nil {
				return filename, file, nil // Success!
			}
		}
	}

	return "", nil, fmt.Errorf("no available filename found in sequence for dir %q with name %q", dir, name)
}