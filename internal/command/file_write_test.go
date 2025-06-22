// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"os"
	"path/filepath"
	"testing"
	"vinti/internal/core"
)

func TestFileWrite_Success(t *testing.T) {
	config := core.DefaultConfig()
	dir := "FileWriteSuccess"
	filename := "testfile.txt"
	data := "hello, Vinti!"

	testPath := filepath.Join(config.Dir, dir)
	os.RemoveAll(testPath)

	// Create the directory (FileWrite expects it to already exist)
	if err := os.MkdirAll(testPath, config.FileModeDir); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Attempt to store the file
	err := FileWrite(config, dir, filename, data)
	if err != nil {
		t.Fatalf("Expected no error on successful write, got: %v", err)
	}

	// Verify contents
	fullPath := filepath.Join(config.Dir, dir, filename)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("Expected to read back written file, got error: %v", err)
	}
	if string(content) != data {
		t.Errorf("Expected file content %q, got %q", data, string(content))
	}
}

func TestFileWrite_Fail(t *testing.T) {
	config := core.DefaultConfig()
	dir := "FileWriteFail" // Do NOT create this directory
	filename := "testfile.txt"

	// Attempt to store the file — should fail since dir doesn't exist
	err := FileWrite(config, dir, filename, "some data")
	if err == nil {
		t.Fatal("Expected error when writing to non-existent directory, got nil")
	}
}
