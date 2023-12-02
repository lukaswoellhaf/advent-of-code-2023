package tools

import "strings"

func SplitInputByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func SplitInputBySemicolon(input string) []string {
	return strings.Split(input, ";")
}

func SplitInputByComma(input string) []string {
	return strings.Split(input, ",")
}
