// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"time"
	"errors"
	"vinti/internal/core"
)

func DataInsert(config *core.Config, dir string, data string) (string, error) {

	baseName := Datetime(time.Now())
	
	filename, file, err := IncrementFile(config, dir, baseName)
	
	if err == nil {
		_, errWrite := file.WriteString(data)
		errClose := file.Close()
		err = errors.Join(errWrite, errClose)
		if err != nil {
			filename = ""
		}
	}
			
	return filename, err
}
