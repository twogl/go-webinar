package ignore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadGitignore() ([]byte, error) {
	absPath, err := filepath.Abs("../.gitignore")
	if err != nil {
		return nil, fmt.Errorf("get absolute file path: %w", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("os open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io read whole file: %w", err)
	}
	return bytes, nil
}
