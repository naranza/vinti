// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "path/filepath"
  "testing"
  "vinti/internal/core"
)

func TestFolderInsert_Success(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "Success"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.RemoveAll(testPath)
  
  err := FolderInsert(config, tmpDir)
  if err != nil {
    t.Fatalf("FolderInsert failed: %v", err)
  }

  info, err := os.Stat(testPath)
  if err != nil {
    t.Fatalf("Expected folder not found: %v", err)
  }
  if !info.IsDir() {
    t.Fatalf("Expected a directory, got something else at: %s", testPath)
  }
}

func TestFolderInsert_Fail(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "Fail"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.RemoveAll(testPath)
  
  file, err := os.Create(testPath)
  file.Close()
  err = FolderInsert(config, tmpDir)

  if err == nil {
    t.Log(testPath)
    t.Log(err)
    t.Fatal("Expected error when trying to create a directory inside a file, got nil")
  }
}
