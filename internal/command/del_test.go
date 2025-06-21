// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"os"
	"path/filepath"
	"testing"

	"vinti/internal/core"
)

func TestDel_Success(t *testing.T) {
	config := core.DefaultConfig()
	dir := "DelSuccess"
	fileName := "testfile.txt"
	content := "delete me"

	// Prepare test directory and file
	testPath := filepath.Join(config.Dir, dir)
	os.RemoveAll(testPath)
	if err := os.MkdirAll(testPath, config.FileModeDir); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	fullPath := filepath.Join(testPath, fileName)
	if err := os.WriteFile(fullPath, []byte(content), config.FileModeFile); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Delete the file
	err := Del(config, dir, fileName)
	if err != nil {
		t.Fatalf("Expected no error deleting file, got %v", err)
	}

	// Verify file no longer exists
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists or error was: %v", err)
	}
}

func TestDel_FileNotExist(t *testing.T) {
	config := core.DefaultConfig()
	dir := "DelFail"
	fileName := "nonexistent.txt"

	// Prepare test directory without the file
	testPath := filepath.Join(config.Dir, dir)
	os.RemoveAll(testPath)
	if err := os.MkdirAll(testPath, config.FileModeDir); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Attempt to delete non-existent file
	err := Del(config, dir, fileName)
	if err == nil {
		t.Fatal("Expected error when deleting non-existent file, got nil")
	}
}
