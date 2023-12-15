package string

import "unicode"

// CameCase Converts string to camel case.
func CamelCase(source string) string {
	res := make([]byte, 0, len(source))
	for i := 0; i < len(source); i++ {
		if unicode.IsLetter(source[i]) {
		}
	}
}
