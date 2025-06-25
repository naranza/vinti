// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "os"
  "time"
  "errors"
  "path/filepath"
  "strings"
  "testing"
  "vinti/internal/core"
)

func DataInsertTest(config *core.Config, dir string, data string, testCase string) (filename string, err error) {

  baseName := Datetime(time.Now())
  
  filename, file, err := IncrementFile(config, dir, baseName)
  
  if err == nil {
    var errWrite error
    if testCase == "write" {
      errWrite = errors.New("write error")
    } else {
      _, errWrite = file.WriteString(data)
    }
    if testCase == "close" {
      file.Close()
    }
    errClose := file.Close()
    err = errors.Join(errWrite, errClose)
    if err != nil {
      filename = ""
    }
  }
      
  return filename, err
}


func TestDataInsert_Success(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "DataInsertTest"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.MkdirAll(testPath, config.FileModeDir)

  data := "hello world"
  
  filename, err := DataInsert(config, tmpDir, data)
  if err != nil {
    t.Fatalf("Add returned error: %v", err)
  }

  // Check filename is not empty and looks like expected format (optional)
  if filename == "" {
    t.Fatal("filename is empty")
  }

  fullPath := filepath.Join(config.Dir, tmpDir, filename)
  
  // Check file exists
  info, err := os.Stat(fullPath)
  if err != nil {
    t.Fatalf("file %q does not exist: %v", fullPath, err)
  }

  // Check file permissions - you can customize this depending on your function behavior
  if info.Mode().Perm() != config.FileModeDir {
    t.Errorf("expected perm %v, got %v", config.FileModeDir, info.Mode().Perm())
  }

  // Check file content
  contentBytes, err := os.ReadFile(fullPath)
  if err != nil {
    t.Fatalf("failed to read file %q: %v", fullPath, err)
  }
  content := string(contentBytes)
  if !strings.Contains(content, data) {
    t.Errorf("file content %q does not contain expected data %q", content, data)
  }
}


func TestDataInsert_WriteError(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "DataInsertTest"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.MkdirAll(testPath, config.FileModeDir)

  data := "hello world"
  
  _, err := DataInsertTest(config, tmpDir, data, "write")
  if err == nil {
    t.Fatalf("Error expected, got: %v", err)
  }
}

func TestDataInsert_CloseError(t *testing.T) {
  config := core.DefaultConfig()
  tmpDir := "DataInsertTest"
  testPath := filepath.Join(config.Dir, tmpDir)
  os.MkdirAll(testPath, config.FileModeDir)

  data := "hello world"
  
  _, err := DataInsertTest(config, tmpDir, data, "close")
  if err == nil {
    t.Fatalf("Error expected, got: %v", err)
  }
}
