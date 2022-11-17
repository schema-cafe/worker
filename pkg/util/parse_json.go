package util

import (
	"encoding/json"
	"strings"
)

func ParseJSON[T any](s string) T {
	var res T
	err := json.NewDecoder(strings.NewReader(s)).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res
}
