package tools

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filepath string) (string, error) {
	content, readErr := os.ReadFile(filepath)
	if readErr != nil {
		log.Fatal(readErr)
		return "", readErr
	}

	return string(content), nil
}

func ReadFileAndTrimSpace(filepath string) (string, error) {
	content, readErr := os.ReadFile(filepath)
	if readErr != nil {
		log.Fatal(readErr)
		return "", readErr
	}

	return strings.TrimSpace(string(content)), nil
}
