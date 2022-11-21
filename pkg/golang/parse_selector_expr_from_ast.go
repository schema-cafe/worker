package golang

import (
	"fmt"
	goast "go/ast"

	"github.com/schema-cafe/go-types/ast"
)

func ParseSelectorExprFromAST(pkgName string, imports []*goast.ImportSpec, exp *goast.SelectorExpr) (ast.Identifier, error) {
	switch x := exp.X.(type) {
	case *goast.Ident:
		pkg, err := ResolveImport(pkgName, imports, x.Name)
		if err != nil {
			return ast.Identifier{}, err
		}
		return ast.Identifier{
			Path: pkg,
			Name: exp.Sel.Name,
		}, nil
	default:
		return ast.Identifier{}, fmt.Errorf("selector expression must be identifier %T", x)
	}
}
