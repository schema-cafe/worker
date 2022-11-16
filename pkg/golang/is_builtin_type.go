package golang

func IsBuiltinType(name string) bool {
	var builtins = []string{
		"bool",
		"string",
		"int",
	}
	for _, t := range builtins {
		if t == name {
			return true
		}
	}
	return false
}
