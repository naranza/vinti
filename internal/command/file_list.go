// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"os"
	"path/filepath"
	"vinti/internal/core"
)

func FileList(config *core.Config, dir string) ([]string, error) {
	fullPath := filepath.Join(config.Dir, dir)

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}