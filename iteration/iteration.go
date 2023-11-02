package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character
	}
	return
}

func RepeatStd(character string, times int) string {
	return strings.Repeat(character, times)
}

func RepeatBuilder(character string, times int) string {
	var b strings.Builder
	for i := 0; i < times; i++ {
		fmt.Fprintf(&b, "%s", character)
	}
	return b.String()
}
