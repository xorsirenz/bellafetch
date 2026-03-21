package utils

import (
	"fmt"
	"os"
)
func OpenFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", filename, err)
	}
	defer file.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %w", filename, err)
	}
	return string(data), nil
}
