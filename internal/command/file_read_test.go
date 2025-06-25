// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "testing"

  "vinti/internal/core"
)

func TestFileRead_Success(t *testing.T) {
  config := core.DefaultConfig()
  dir := "GetSuccess"
  fileName := "testfile.txt"
  content := "hello vinti"

  // Prepare test directory
  testPath := filepath.Join(config.Dir, dir)
  os.RemoveAll(testPath)
  if err := os.MkdirAll(testPath, config.FileModeDir); err != nil {
    t.Fatalf("Failed to create test directory: %v", err)
  }

  fullPath := filepath.Join(testPath, fileName)
  if err := os.WriteFile(fullPath, []byte(content), config.FileModeFile); err != nil {
    t.Fatalf("Failed to write test file: %v", err)
  }

  result, err := FileRead(config, dir, fileName)
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }
  if result != content {
    t.Errorf("Expected content %q, got %q", content, result)
  }
}

func TestFileRead_FileNotExist(t *testing.T) {
  config := core.DefaultConfig()
  dir := "GetFail"
  fileName := "nonexistent.txt"

  // Make sure the file does not exist
  testPath := filepath.Join(config.Dir, dir)
  os.RemoveAll(testPath)
  os.MkdirAll(testPath, config.FileModeDir)

  _, err := FileRead(config, dir, fileName)
  if err == nil {
    t.Fatal("Expected error when getting non-existent file, got nil")
  }
}
