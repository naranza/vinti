// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"os"
	"path/filepath"
	"testing"
	"vinti/internal/core"
)

func TestAll_Success(t *testing.T) {
	config := core.DefaultConfig()
	dir := "AllSuccess"
	testPath := filepath.Join(config.Dir, dir)
	os.RemoveAll(testPath)

	// Create test directory and files
	if err := os.MkdirAll(testPath, config.FileModeDir); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	filesToCreate := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, f := range filesToCreate {
		fullPath := filepath.Join(testPath, f)
		if err := os.WriteFile(fullPath, []byte("content"), config.FileModeFile); err != nil {
			t.Fatalf("Failed to create test file %s: %v", f, err)
		}
	}

	files, err := All(config, dir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if returned files match what was created
	if len(files) != len(filesToCreate) {
		t.Errorf("Expected %d files, got %d", len(filesToCreate), len(files))
	}

	fileSet := make(map[string]bool)
	for _, f := range files {
		fileSet[f] = true
	}
	for _, f := range filesToCreate {
		if !fileSet[f] {
			t.Errorf("Expected file %q to be listed but it was not", f)
		}
	}
}

func TestAll_Fail(t *testing.T) {
	config := core.DefaultConfig()
	dir := "AllFail"
	testPath := filepath.Join(config.Dir, dir)
	os.RemoveAll(testPath)

	// Don't create directory, so it doesn't exist

	_, err := All(config, dir)
	if err == nil {
		t.Fatal("Expected error when listing files from non-existent directory, got nil")
	}
}
