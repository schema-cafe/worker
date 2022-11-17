package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/schema-cafe/go-types"
)

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
