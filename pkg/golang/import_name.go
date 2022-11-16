package golang

import (
	"go/ast"
	"strconv"
	"strings"
)

func ImportName(imp *ast.ImportSpec) string {
	if imp.Name != nil {
		return imp.Name.Name
	} else {
		path, _ := strconv.Unquote(imp.Path.Value)
		parts := strings.Split(path, "/")
		n := parts[len(parts)-1]
		if strings.HasPrefix(n, "go-") {
			n = strings.TrimPrefix(n, "go-")
		}
		return n
	}
}
