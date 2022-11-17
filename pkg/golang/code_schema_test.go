package golang_test

import (
	"testing"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/worker/pkg/golang"
)

func TestCodeSchema(t *testing.T) {
	s := types.Schema{
		Fields: []types.Field{
			{
				Name: "ID",
				Type: types.Type{
					BaseType: types.Identifier{
						Path: "github.com/schema-cafe/go-types",
						Name: "ID",
					},
				},
			},
			{
				Name: "Name",
				Type: types.Type{
					BaseType: types.Identifier{
						Path: " ",
						Name: "string",
					},
				},
			},
			{
				Name: "Age",
				Type: types.Type{
					BaseType: types.Identifier{
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
