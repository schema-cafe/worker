package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := getEnv("SCHEMA_CAFE_ORG_DIR")
	checkError(err)
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
