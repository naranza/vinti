// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "testing"

  "vinti/internal/core"
)

func TestFileRename_Success(t *testing.T) {
  config := core.DefaultConfig()
  base := "TestFileRename_Success"
  srcDir := filepath.Join(base, "source")
  dstDir := filepath.Join(base, "dest")
  filename := "file.txt"
  content := "test"

  os.RemoveAll(filepath.Join(config.Dir, base))
	
	os.MkdirAll(filepath.Join(config.Dir, srcDir), config.FileModeDir)
	os.MkdirAll(filepath.Join(config.Dir, dstDir), config.FileModeDir)

  srcPath := filepath.Join(config.Dir, srcDir, filename)
	dstPath := filepath.Join(config.Dir, dstDir, filename)
	os.WriteFile(srcPath, []byte(content), config.FileModeFile)

	// Perform rename
   err := FileRename(config, srcDir, filename, dstDir)
  if err != nil {
    t.Fatalf("Expected no error renaming file, got: %v", err)
  }

  // Verify old file no longer exists
	_, err = os.Stat(srcPath)
	if !os.IsNotExist(err) {
    t.Fatalf("Expected source file to be gone, still exists or error: %v", err)
  }

  // Verify new file exists
	 _, err = os.Stat(dstPath)
  if; err != nil {
    t.Fatalf("Expected new file to exist, got error: %v", err)
  }
}

func TestFileRename_FileNotExist(t *testing.T) {
  config := core.DefaultConfig()
  base := "TestFileRename_FileNotExist"
  srcDir := filepath.Join(base, "source")
  dstDir := filepath.Join(base, "dest")
  filename := "file.txt"
  content := "test"

  os.RemoveAll(filepath.Join(config.Dir, base))
	
	os.MkdirAll(filepath.Join(config.Dir, srcDir), config.FileModeDir)
	os.MkdirAll(filepath.Join(config.Dir, dstDir), config.FileModeDir)

  srcPath := filepath.Join(config.Dir, srcDir, filename)
	os.WriteFile(srcPath, []byte(content), config.FileModeFile)

  // Perform rename
  err := FileRename(config, srcDir, "123", dstDir)
  if err == nil {
    t.Fatal("Expected error when renaming non-existent file, got nil")
  }
}



