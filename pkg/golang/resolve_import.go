package golang

import (
	"fmt"
	"go/ast"
	"strconv"
)

func ResolveImport(pkgName string, imports []*ast.ImportSpec, name string) (string, error) {
	for _, imp := range imports {
		if name == ImportName(imp) {
			path, _ := strconv.Unquote(imp.Path.Value)
			return path, nil
		}
	}
	return "", fmt.Errorf("import not found: %s", name)
}
