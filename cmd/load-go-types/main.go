package main

import "github.com/schema-cafe/worker/pkg/golang"

func main() {
	const goTypesRepo = "../go-types"
	const target = "../../../../data/schema.cafe/public/schemas"
	copy(goTypesRepo, target, []string{})
}

func copy(from string, to string, at []string) {
	r := golang.TypesRepo{
		Dir: from,
	}
	node, err := r.GetNode(at)
	if err != nil {
		panic(err)
	}
	if node.IsFolder {
		for _, child := range node.Folder.Contents {
			copy(from, to, append(at, child.Name))
		}
	} else {
		schema := node.Value
	}
}
