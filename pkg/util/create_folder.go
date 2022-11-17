package util

import (
	"os"
	"path/filepath"
)

func CreateFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(path, ".keep"), []byte{}, os.ModePerm)
}
