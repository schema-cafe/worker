package queries

import (
	"github.com/schema-cafe/worker/pkg/golang"
	"github.com/schema-cafe/worker/pkg/web"
)

func GetNode(workdir string, inputs map[string]string) (any, error) {
	repo := golang.TypesRepo{
		Dir: workdir,
	}
	return repo.GetNode(web.Path(inputs["path"]))
}
