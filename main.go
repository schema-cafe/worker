package main

import (
	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/worker/pkg/commands"
	"github.com/schema-cafe/worker/pkg/queries"
	"github.com/schema-cafe/worker/pkg/util"
	"github.com/schema-cafe/worker/pkg/web"
)

func main() {
	workdir, err := util.GetEnv("WORKDIR")
	util.PanicIfError(err)

	port, err := util.GetEnv("PORT")
	util.PanicIfError(err)

	app := types.API{
		Queries: map[string]types.QueryFunction{
			"/": queries.GetNode,
		},
		Commands: map[string]types.CommandFunction{
			"/cmd/create-folder":         commands.CreateFolder,
			"/cmd/move":                  commands.Move,
			"/cmd/delete":                commands.Delete,
			"/cmd/create-schema":         commands.CreateSchema,
			"/cmd/add-schema-field":      commands.AddSchemaField,
			"/cmd/rename-schema-field":   commands.RenameSchemaField,
			"/cmd/set-schema-field-type": commands.SetSchemaFieldType,
			"/cmd/delete-schema-field":   commands.DeleteSchemaField,
		},
	}

	err = web.ServeAPI(app, workdir, port)
	util.PanicIfError(err)
}
