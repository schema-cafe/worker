package golang

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/go-types/filesystem"
)

type TypesRepo struct {
	Dir string

	pkgName string
}

func (r *TypesRepo) GetNode(path []string) (*filesystem.Node[types.Schema], error) {
	pkgName, err := r.PackageName(path)
	if err != nil {
		return nil, err
	}
	p := r.getPath(path)
	gofilepath := p + ".go"
	if fileExists(gofilepath) {
		s, err := ParseSchemaFromFile(pkgName, gofilepath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse schema: %w", err)
		}
		return &filesystem.Node[types.Schema]{
			IsFolder: false,
			Value:    s,
		}, nil
	} else {
		folder, err := ReadDir(p)
		if err != nil {
			return nil, err
		}
		return &filesystem.Node[types.Schema]{
			IsFolder: true,
			Folder:   folder,
		}, nil
	}
}

func (r *TypesRepo) PackageName(path []string) (string, error) {
	err := r.loadPkgName()
	if err != nil {
		return "", err
	}
	return filepath.Join(r.pkgName, filepath.Join(path[:len(path)-1]...)), nil
}

func (r *TypesRepo) getPath(path []string) string {
	return filepath.Join(r.Dir, filepath.Join(path...))
}

func (r *TypesRepo) loadPkgName() error {
	if r.pkgName != "" {
		return nil
	}
	path := filepath.Join(r.Dir, "go.mod")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")
	if len(lines) == 0 {
		return fmt.Errorf("go.mod is empty")
	}
	r.pkgName = strings.TrimPrefix(lines[0], "module ")
	err = ValidatePackageName(r.pkgName)
	if err != nil {
		return fmt.Errorf("invalid package name in go.mod: %w", err)
	}
	return nil
}

func ensure(err error) {
	if err != nil {
		panic(err)
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
