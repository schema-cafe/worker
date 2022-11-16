package golang

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/schema-cafe/go-types"
)

func ParseSchemaFromFile(pkgName, gofilepath string) (*types.Schema, error) {
	s := types.Schema{}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, gofilepath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							return ParseSchemaFromAST(pkgName, f.Imports, structType)
						}
					}
				}
			}
		}
	}
	return &s, nil
}
