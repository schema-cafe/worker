package golang

import (
	goast "go/ast"

	"github.com/library-development/go-nameconv"
	"github.com/library-development/go-schemacafe"
)

func ParseSchemaFromAST(pkgName string, imports []*goast.ImportSpec, s *goast.StructType) (*schemacafe.Schema, error) {
	schema := schemacafe.Schema{}
	for _, field := range s.Fields.List {
		fieldType, err := ParseTypeFromAST(pkgName, imports, field.Type)
		if err != nil {
			return nil, err
		}
		for _, name := range field.Names {
			schema.Fields = append(schema.Fields, schemacafe.Field{
				Name: nameconv.ParsePascalCase(name.Name),
				Type: fieldType,
			})
		}
	}
	return &schema, nil
}
