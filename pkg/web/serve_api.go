package web

import (
	"net/http"

	"github.com/schema-cafe/go-types"
)

func ServeAPI(app types.API, workdir, port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleCORS(w, r)
		switch r.Method {
		case "GET":
			q, ok := app.Queries[r.URL.Path]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			HandleQuery(workdir, q, w, r)
		case "POST":
			c, ok := app.Commands[r.URL.Path]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			HandleCommand(workdir, c, w, r)
		}
	})
	return http.ListenAndServe(":"+port, nil)
}
