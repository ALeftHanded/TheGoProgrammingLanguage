package string_concatenation

import (
	"fmt"
	"strings"
)

// ConcatenateStringsWithPlus concatenates strings with plus operator.
func ConcatenateStringsWithPlus(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}

// ConcatenateStringsWithBuilder concatenates strings with strings.Builder.
func ConcatenateStringsWithBuilder(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

// ConcatenateStringsWithJoin concatenates strings with strings.Join.
func ConcatenateStringsWithJoin(strs []string) string {
	return strings.Join(strs, "")
}

// ConcatenateStringsWithSprintf concatenates strings with fmt.Sprintf.
func ConcatenateStringsWithSprintf(iStrs []interface{}) string {
	formatString := strings.Repeat("%s", len(iStrs))
	return fmt.Sprintf(formatString, iStrs...)
}

// ConcatenateStringsWithBuilderAndGrow concatenates strings with strings.Builder and builder grow.
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
