// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package api

import (
	"regexp"
	"strings"
)

var folderRegex = regexp.MustCompile(`^[a-zA-Z0-9/]+$`)
var fileRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]+(\.[a-zA-Z0-9]+)?$`)

func ValidateFolder(folder string) bool {
	if !folderRegex.MatchString(folder) {
		return false
	}
	if strings.HasPrefix(folder, "/") || strings.HasSuffix(folder, "/") {
		return false
	}
	if strings.Contains(folder, "..") || strings.Contains(folder, "//") {
		return false
	}
	return true
}

func ValidateFile(file string) bool {
	return fileRegex.MatchString(file)
}