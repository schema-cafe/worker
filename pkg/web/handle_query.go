package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/schema-cafe/go-types"
)

func HandleQuery(workdir string, q types.QueryFunction, w http.ResponseWriter, r *http.Request) {
	inputs := map[string]string{}
	for key, values := range r.URL.Query() {
		inputs[key] = values[0]
	}
	result, err := q(workdir, inputs)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(types.QueryResult{
		Result: result,
		Error:  err,
	})
}
