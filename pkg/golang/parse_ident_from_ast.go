package golang

import (
	goast "go/ast"

	"github.com/schema-cafe/go-types/ast"
)

func ParseIdentFromAST(pkgName string, ident *goast.Ident) (ast.Identifier, error) {
	name := ident.Name
	id := ast.Identifier{
		Name: name,
	}
	if !IsBuiltinType(name) {
		id.Path = pkgName
	}
	return id, nil
}
