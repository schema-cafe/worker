package golang_test

import (
	"testing"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/go-types/ast"
	"github.com/schema-cafe/go-types/form"
	"github.com/schema-cafe/worker/pkg/golang"
)

func TestCodeSchema(t *testing.T) {
	s := types.Schema{
		Fields: []form.Field{
			{
				Name: "ID",
				Type: ast.Type{
					BaseType: ast.Identifier{
						Path: "github.com/schema-cafe/go-types",
						Name: "ID",
					},
				},
			},
			{
				Name: "Name",
				Type: ast.Type{
					BaseType: ast.Identifier{
						Path: " ",
						Name: "string",
					},
				},
			},
			{
				Name: "Age",
				Type: ast.Type{
					BaseType: ast.Identifier{
						Path: "github.com/schema-cafe/go-types/people",
						Name: "Age",
					},
				},
			},
		},
	}

	path := "github.com/schema-cafe/go-types"
	code := golang.CodeSchema(path, &s)

	expected := `package types

import (
	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/go-types/people"
)

type types struct {
	ID types.ID
	Name string
	Age people.Age
}`

	if code != expected {
		t.Errorf("expected %s, got %s", expected, code)
	}
}
