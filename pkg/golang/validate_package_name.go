package golang

import (
	"fmt"
	"strings"
)

func ValidatePackageName(pkgName string) error {
	if pkgName == "" {
		return fmt.Errorf("empty")
	}
	if strings.Contains(pkgName, " ") {
		return fmt.Errorf("contains spaces")
	}
	return nil
}
