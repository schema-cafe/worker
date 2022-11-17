package golang

import (
	"path/filepath"
	"strings"

	"github.com/schema-cafe/go-types"
)

func CodeSchema(path string, schema *types.Schema) string {
	var typeString strings.Builder
	imports := map[string]bool{}

	typeString.WriteString("type ")
	typeString.WriteString(filepath.Base(path))
	typeString.WriteString(" struct {\n")
	for _, field := range schema.Fields {
		typeString.WriteString("\t")
		typeString.WriteString(field.Name)
		typeString.WriteString(" ")
		if field.Type.IsArray {
			typeString.WriteString("[]")
		}
		if field.Type.IsMap {
			typeString.WriteString("map[string]")
		}
		if field.Type.IsPointer {
			typeString.WriteString("*")
		}
		if field.Type.BaseType.Path != filepath.Dir(path) {
			imports[field.Type.BaseType.Path] = true
			typeString.WriteString(filepath.Base(filepath.Dir(field.Type.BaseType.Path)))
			typeString.WriteString(".")
		}
		typeString.WriteString(field.Type.BaseType.Name)
		typeString.WriteString("\n")
	}
	typeString.WriteString("}\n")

	var s strings.Builder
	s.WriteString("package ")
	s.WriteString(filepath.Base(filepath.Dir(path)))
	s.WriteString("\n\n")
	if len(imports) > 0 {
		s.WriteString("import (\n")
		for path := range imports {
			s.WriteString("\t\"")
			s.WriteString(path)
			s.WriteString("\"\n")
		}
		s.WriteString(")\n\n")
	}

	s.WriteString(typeString.String())

	return s.String()
}
