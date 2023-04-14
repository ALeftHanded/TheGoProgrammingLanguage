package string_concatenation

import "strings"

func ConcatenateStrings(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}

func ConcatenateStringsOp(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
