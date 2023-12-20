package string

import (
	"strings"
)

// CapitalizeTitle
// Converts the first character of string to upper case and the remaining to lower case.
func CapitalizeTitle(title string) string {
	title = strings.ToLower(title)
	v := strings.Split(title, " ")

	for i := 0; i < len(v); i++ {
		if len(v[i]) < 3 {
			v[i] = strings.ToLower(v[i])
		} else {
			v[i] = strings.Title(v[i])
		}
	}

	title = strings.Join(v, " ")
	return title
}
