package core

import (
	"time"
	"errors"
)

func Add(config *Config, dir string, data string) (string, error) {

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


