package golang

import (
	"fmt"
	"go/ast"

	"github.com/schema-cafe/go-types"
)

func ParseTypeFromAST(pkgName string, imports []*ast.ImportSpec, t ast.Expr) (types.Type, error) {
	switch t := t.(type) {
	case *ast.Ident:
		baseType, err := ParseIdentFromAST(pkgName, t)
		return types.Type{
			BaseType: baseType,
		}, err
	case *ast.SelectorExpr:
		baseType, err := ParseSelectorExprFromAST(pkgName, imports, t)
		return types.Type{
			BaseType: baseType,
		}, err
	case *ast.ArrayType:
		elemType, err := ParseTypeFromAST(pkgName, imports, t.Elt)
		if err != nil {
			return types.Type{}, err
		}
		if elemType.IsArray || elemType.IsMap || elemType.IsPointer {
			return types.Type{}, fmt.Errorf("array element type must not be array, map or pointer")
		}
		return types.Type{
			IsArray:  true,
			BaseType: elemType.BaseType,
		}, err
	case *ast.MapType:
		keyType, err := ParseTypeFromAST(pkgName, imports, t.Key)
		if err != nil {
			return types.Type{}, err
		}
		if keyType.IsArray || keyType.IsMap || keyType.IsPointer {
			return types.Type{}, fmt.Errorf("map key type must not be array, map or pointer")
		}
		if keyType.BaseType.Name != "string" {
			return types.Type{}, fmt.Errorf("map key type must be string")
		}
		valueType, err := ParseTypeFromAST(pkgName, imports, t.Value)
		if err != nil {
			return types.Type{}, err
		}
		if valueType.IsArray || valueType.IsMap || valueType.IsPointer {
			return types.Type{}, fmt.Errorf("map value type must not be array, map or pointer")
		}
		return types.Type{
			IsMap:    true,
			BaseType: valueType.BaseType,
		}, err
	case *ast.StarExpr:
		baseType, err := ParseTypeFromAST(pkgName, imports, t.X)
		if err != nil {
			return types.Type{}, err
		}
		if baseType.IsArray || baseType.IsMap || baseType.IsPointer {
			return types.Type{}, fmt.Errorf("pointer type must not be array, map or pointer")
		}
		return types.Type{
			IsPointer: true,
			BaseType:  baseType.BaseType,
		}, err
	default:
		return types.Type{}, fmt.Errorf("unsupported type: %T", t)
	}
}
