package main

import (
	"os"
	"path/filepath"
	"strings"
)

func getExtension(filePath string) string {
	filePathParts := strings.Split(strings.TrimSpace(filePath), "/")
	fileName := filePathParts[len(filePathParts)-1]
	fileNameElems := strings.Split(strings.TrimSpace(fileName), ".")
	extension := strings.TrimPrefix(fileName, fileNameElems[0])
	return extension
}

func createDirectoryPath(path string) {
	err := os.MkdirAll(filepath.Dir(path), 0770)
	checkError(err)
}
