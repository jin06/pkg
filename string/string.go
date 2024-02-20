package string

import (
	"strings"
)

// KebabCase Converts string to kebab case
func KebabCase(s string) string {
	items := splitLetter(s)
	for i := 0; i < len(items); i++ {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "-")
}

func Pad(s string, i int) string {
	return ""
}
