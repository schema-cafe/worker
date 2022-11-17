package commands

import (
	"fmt"
	"path/filepath"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/worker/pkg/util"
)

func CreateFolder(stateDir string, inputs map[string]string) ([]types.Mutation, error) {
	in := inputs["in"]
	name := inputs["name"]

	if !util.IsFolder(filepath.Join(stateDir, in)) {
		return nil, fmt.Errorf("%v is not a folder", in)
	}

	path := filepath.Join(stateDir, in, name)
	if util.Exists(path) || util.Exists(path+".go") {
		return nil, fmt.Errorf("%v already exists", path)
	}

	return []types.Mutation{
		{
			Path:     filepath.Join(path, ".keep"),
			Contents: " ",
		},
	}, nil
}
