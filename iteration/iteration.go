package iteration

import "strings"

func Repeat(character string, length int) string {
	var repeated strings.Builder
	for range length {
		repeated.WriteString(character)
	}

	return repeated.String()
}
