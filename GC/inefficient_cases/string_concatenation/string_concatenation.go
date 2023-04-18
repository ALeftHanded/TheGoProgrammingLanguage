package string_concatenation

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}

func ConcatenateStringsWithBuilder(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

func ConcatenateStringsWithJoin(strs []string) string {
	return strings.Join(strs, "")
}

func ConcatenateStringsWithSprintf() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
}

func ConcatenateStringsWithBuilderAndGrow(strs []string) string {
	var builder strings.Builder
	totalLength := 0
	for _, str := range strs {
		totalLength += len(str)
	}
	builder.Grow(totalLength)
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
