package util

import (
	"bytes"
	"encoding/json"
)

func EncodeJSON(s any) string {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(s)
	if err != nil {
		panic(err)
	}
	return b.String()
}
