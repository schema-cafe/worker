package golang

import (
	"fmt"
	goast "go/ast"

	"github.com/schema-cafe/go-types/ast"
)

func ParseTypeFromAST(pkgName string, imports []*goast.ImportSpec, t goast.Expr) (ast.Type, error) {
	switch t := t.(type) {
	case *goast.Ident:
		baseType, err := ParseIdentFromAST(pkgName, t)
		return ast.Type{
			BaseType: baseType,
		}, err
	case *goast.SelectorExpr:
		baseType, err := ParseSelectorExprFromAST(pkgName, imports, t)
		return ast.Type{
			BaseType: baseType,
		}, err
	case *goast.ArrayType:
		elemType, err := ParseTypeFromAST(pkgName, imports, t.Elt)
		if err != nil {
			return ast.Type{}, err
		}
		if elemType.IsArray || elemType.IsMap || elemType.IsPointer {
			return ast.Type{}, fmt.Errorf("array element type must not be array, map or pointer")
		}
		return ast.Type{
			IsArray:  true,
			BaseType: elemType.BaseType,
		}, err
	case *goast.MapType:
		keyType, err := ParseTypeFromAST(pkgName, imports, t.Key)
		if err != nil {
			return ast.Type{}, err
		}
		if keyType.IsArray || keyType.IsMap || keyType.IsPointer {
			return ast.Type{}, fmt.Errorf("map key type must not be array, map or pointer")
		}
		if keyType.BaseType.Name != "string" {
			return ast.Type{}, fmt.Errorf("map key type must be string")
		}
		valueType, err := ParseTypeFromAST(pkgName, imports, t.Value)
		if err != nil {
			return ast.Type{}, err
		}
		if valueType.IsArray || valueType.IsMap || valueType.IsPointer {
			return ast.Type{}, fmt.Errorf("map value type must not be array, map or pointer")
		}
		return ast.Type{
			IsMap:    true,
			BaseType: valueType.BaseType,
		}, err
	case *goast.StarExpr:
		baseType, err := ParseTypeFromAST(pkgName, imports, t.X)
		if err != nil {
			return ast.Type{}, err
		}
		if baseType.IsArray || baseType.IsMap || baseType.IsPointer {
			return ast.Type{}, fmt.Errorf("pointer type must not be array, map or pointer")
		}
		return ast.Type{
			IsPointer: true,
			BaseType:  baseType.BaseType,
		}, err
	default:
		return ast.Type{}, fmt.Errorf("unsupported type: %T", t)
	}
}
