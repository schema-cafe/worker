package util

import (
	"os"
	"path/filepath"

	"github.com/schema-cafe/go-types"
)

func ApplyMutations(dir string, mutations []types.Mutation) error {
	for _, m := range mutations {
		path := filepath.Join(dir, m.Path)
		if m.Contents == "" {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		} else {
			contents := []byte{}
			if m.Contents != " " {
				contents = []byte(m.Contents)
			}
			err := os.MkdirAll(filepath.Dir(path), 0755)
			if err != nil {
				return err
			}
			err = os.WriteFile(path, contents, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
