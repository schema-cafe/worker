package golang

import (
	"fmt"
	"go/ast"

	"github.com/schema-cafe/go-types"
)

func ParseSelectorExprFromAST(pkgName string, imports []*ast.ImportSpec, exp *ast.SelectorExpr) (types.Identifier, error) {
	switch x := exp.X.(type) {
	case *ast.Ident:
		pkg, err := ResolveImport(pkgName, imports, x.Name)
		if err != nil {
			return types.Identifier{}, err
		}
		return types.Identifier{
			Pkg:  pkg,
			Name: exp.Sel.Name,
		}, nil
	default:
		return types.Identifier{}, fmt.Errorf("selector expression must be identifier %T", x)
	}
}
