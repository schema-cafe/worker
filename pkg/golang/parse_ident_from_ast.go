package golang

import (
	"go/ast"

	"github.com/schema-cafe/go-types"
)

func ParseIdentFromAST(pkgName string, ident *ast.Ident) (types.Identifier, error) {
	name := ident.Name
	id := types.Identifier{
		Name: name,
	}
	if !IsBuiltinType(name) {
		id.Path = pkgName
	}
	return id, nil
}
