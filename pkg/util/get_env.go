package util

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%v is not set", key)
	}
	return value, nil
}
