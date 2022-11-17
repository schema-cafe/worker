package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/schema-cafe/go-types"
	"github.com/schema-cafe/worker/pkg/util"
)

func HandleCommand(goTypesDir string, c types.CommandFunction, w http.ResponseWriter, r *http.Request) {
	inputs := map[string]string{}
	json.NewDecoder(r.Body).Decode(&inputs)
	mutations, err := c(goTypesDir, inputs)
	if err != nil {
		fmt.Println(err)
	} else {
		err = util.ApplyMutations(goTypesDir, mutations)
		if err != nil {
			fmt.Println(err)
			// TODO: Rollback any partial changes
		}
	}
	json.NewEncoder(w).Encode(types.CommandResult{
		Mutations: mutations,
		Error:     err,
	})
}
