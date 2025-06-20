package core

import (
	"os"
	"time"
	"errors"
	"path/filepath"
	"strings"
	"testing"
)

func AddTest(config *Config, dir string, data string, testCase string) (filename string, err error) {

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


func TestVintiAdd_Success(t *testing.T) {
	config := DefaultConfig()
	tmpDir := "AddTest"
	testPath := filepath.Join(config.Dir, tmpDir)
	os.MkdirAll(testPath, config.FileModeDir)

	data := "hello world"
	
	filename, err := Add(config, tmpDir, data)
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


func TestVintiAdd_WriteError(t *testing.T) {
	config := DefaultConfig()
	tmpDir := "AddTest"
	testPath := filepath.Join(config.Dir, tmpDir)
	os.MkdirAll(testPath, config.FileModeDir)

	data := "hello world"
	
	_, err := AddTest(config, tmpDir, data, "write")
	if err == nil {
		t.Fatalf("Error expected, got: %v", err)
	}
}

func TestVintiAdd_CloseError(t *testing.T) {
	config := DefaultConfig()
	tmpDir := "AddTest"
	testPath := filepath.Join(config.Dir, tmpDir)
	os.MkdirAll(testPath, config.FileModeDir)

	data := "hello world"
	
	_, err := AddTest(config, tmpDir, data, "close")
	if err == nil {
		t.Fatalf("Error expected, got: %v", err)
	}
}

// You can add more tests for failure cases if needed, e.g., invalid dir, no permission, etc.
