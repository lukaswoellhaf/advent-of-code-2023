package tools

import "strings"

func SplitInputByNewLine(input string) []string {
	return strings.Split(input, "\n")
}
