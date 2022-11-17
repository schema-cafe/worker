package git

import "strings"

func GenerateCommitMessage(cmd string, inputs map[string]string) string {
	var s strings.Builder
	s.WriteString(cmd)
	s.WriteString("\n")
	if len(inputs) > 0 {
		s.WriteString("\n")
	}
	for k, v := range inputs {
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v)
		s.WriteString("\n")
	}
	return s.String()
}
