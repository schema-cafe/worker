package golang

import (
	goast "go/ast"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/go-types/form"
)

func ParseSchemaFromAST(pkgName string, imports []*goast.ImportSpec, s *goast.StructType) (*types.Schema, error) {
	schema := types.Schema{}
	for _, field := range s.Fields.List {
		fieldType, err := ParseTypeFromAST(pkgName, imports, field.Type)
		if err != nil {
			return nil, err
		}
		for _, name := range field.Names {
			schema.Fields = append(schema.Fields, form.Field{
				Name: name.Name,
				Type: fieldType,
			})
		}
	}
	return &schema, nil
}
