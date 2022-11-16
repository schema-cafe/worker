package web

import "strings"

func Path(p string) []string {
	path := []string{}
	if p == "" {
		return path
	}
	if p == "/" {
		return path
	}
	for _, part := range strings.Split(p[1:], "/") {
		path = append(path, part)
	}
	return path
}
