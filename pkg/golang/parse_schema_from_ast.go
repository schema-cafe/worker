package golang

import (
	"go/ast"

	"github.com/schema-cafe/go-types"
)

func ParseSchemaFromAST(pkgName string, imports []*ast.ImportSpec, s *ast.StructType) (*types.Schema, error) {
	schema := types.Schema{}
	for _, field := range s.Fields.List {
		fieldType, err := ParseTypeFromAST(pkgName, imports, field.Type)
		if err != nil {
			return nil, err
		}
		for _, name := range field.Names {
			schema.Fields = append(schema.Fields, types.Field{
				Name: name.Name,
				Type: fieldType,
			})
		}
	}
	return &schema, nil
}
