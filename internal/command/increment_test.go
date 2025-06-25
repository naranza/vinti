// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "testing"
  "vinti/internal/core"
)

func TestIncrementFile_FirstAvailable(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "FirstAvailable"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.RemoveAll(testPath)
  os.MkdirAll(testPath, config.FileModeDir)
  
  filename, file, err := IncrementFile(config, tmpDir, "testfile")
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }
  if file == nil {
    t.Fatal("Expected a non-nil file, got nil")
  }
  file.Close()

  expectedFilename := "testfile000000"
  if filename != expectedFilename {
    t.Errorf("Expected filename %q, got %q", expectedFilename, filename)
  }

  fileExpectecd := filepath.Join(config.Dir, tmpDir, expectedFilename)
  if _, err := os.Stat(fileExpectecd); os.IsNotExist(err) {
    t.Errorf("Expected file %q to exist, but it does not", fileExpectecd)
  }
}

func TestIncrementFile_LaterAvailable(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "LaterAvailable"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.RemoveAll(testPath)
  os.MkdirAll(testPath, config.FileModeDir)

  filename, file, err := IncrementFile(config, tmpDir, "testfile")
  file.Close()
  filename, file, err = IncrementFile(config, tmpDir, "testfile")
  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }
  if file == nil {
    t.Fatal("Expected a non-nil file, got nil")
  }
  file.Close()

  expectedFilename := "testfile000001"
  if filename != expectedFilename {
    t.Errorf("Expected filename %q, got %q", expectedFilename, filename)
  }

  fileExpectecd := filepath.Join(config.Dir, tmpDir, expectedFilename)
  if _, err := os.Stat(fileExpectecd); os.IsNotExist(err) {
    t.Errorf("Expected file %q to exist, but it does not", fileExpectecd)
  }
}

func TestIncrementFile_NoAvailable(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "NoAvailable"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.RemoveAll(testPath)
  os.MkdirAll(testPath, config.FileModeDir)
  config.IncrementMax = 1

  _, file, err := IncrementFile(config, tmpDir, "testfile")
  file.Close()
  _, file, err = IncrementFile(config, tmpDir, "testfile")
  file.Close()
  _, file, err = IncrementFile(config, tmpDir, "testfile")

  if err == nil {
    t.Errorf("Expected error message contains: no available filename found in sequence for dir %q with name %q", tmpDir, "testfile")
  }
}

