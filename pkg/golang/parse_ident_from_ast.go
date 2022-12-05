package golang

import (
	goast "go/ast"

	"github.com/library-development/go-schemacafe"
)

func ParseIdentFromAST(pkgPath schemacafe.Path, ident *goast.Ident) (schemacafe.Identifier, error) {
	name := ident.Name
	id := schemacafe.Identifier{
		Path: schemacafe.Path{{pkgName}},
	}
	if !IsBuiltinType(name) {
		id.Path = pkgName
	}
	return id, nil
}
