package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/worker/pkg/queries"
)

func main() {
	workdir, err := getEnv("SCHEMA_CAFE_ORG_DIR")
	checkError(err)

	port, err := getEnv("PORT")
	checkError(err)

	app := types.API{
		Queries: map[string]types.QueryFunction{
			"/": queries.GetNode,
		},
		Commands: map[string]types.CommandFunction{},
	}

	err = ServeAPI(app, workdir, port)
	checkError(err)
}

func ServeAPI(app types.API, workdir, port string) error {
	goTypesDir := filepath.Join(workdir, "go-types")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			q, ok := app.Queries[r.URL.Path]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			HandleQuery(goTypesDir, q, w, r)
		case "POST":
			c, ok := app.Commands[r.URL.Path]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			HandleCommand(goTypesDir, c, w, r)
		}
	})
	return http.ListenAndServe(":"+port, nil)
}

func HandleQuery(goTypesDir string, q types.QueryFunction, w http.ResponseWriter, r *http.Request) {
	inputs := map[string]string{}
	for key, values := range r.URL.Query() {
		inputs[key] = values[0]
	}
	result, err := q(goTypesDir, inputs)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(types.QueryResult{
		Result: result,
		Error:  err,
	})
}

func HandleCommand(goTypesDir string, c types.CommandFunction, w http.ResponseWriter, r *http.Request) {
	inputs := map[string]string{}
	json.NewDecoder(r.Body).Decode(&inputs)
	mutation, err := c(goTypesDir, inputs)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(types.CommandResult{
		Mutation: mutation,
		Error:    err,
	})
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%v is not set", key)
	}
	return value, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
