// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"errors"
	"os"
	"path/filepath"
	"vinti/internal/core"
)

func Sto(config *core.Config, dir string, filename string, data string) error {
	fullPath := filepath.Join(config.Dir, dir, filename)

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, config.FileModeFile)
	if err == nil {
		_, errWrite := file.WriteString(data)
		errClose := file.Close()
		err = errors.Join(errWrite, errClose)
	}
	return err
}
