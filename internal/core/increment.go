package core

import (
  "fmt"
  "os"
  "path/filepath"
)

func IncrementFile(config *Config, dir, name string) (string, *os.File, error) {
	for i := 0; i <= config.IncrementMax; i++ {
		filename := fmt.Sprintf("%s%0*d", name, config.increment_digits, i)
		path := filepath.Join(config.Dir, dir, filename)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			file, err := os.Create(path)
			if err == nil {
				fmt.Println(err)
				fmt.Println(path)
				return filename, file, nil // Success!
			}
		}
	}

	return "", nil, fmt.Errorf("no available filename found in sequence for dir %q with name %q", dir, name)
}