package string

func isLetter(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isUpper(r byte) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r byte) bool {
	return r >= 'a' && r <= 'z'
}

func toUpper(r byte) byte {
	if isLower(r) {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r byte) byte {
	if isUpper(r) {
		return r + ('a' - 'A')
	}
	return r
}

func isLetterRune(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isUpperRune(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLowerRune(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func toUpperRune(r rune) rune {
	if isLowerRune(r) {
		return r - ('a' - 'A')
	}
	return r
}

func toLowerRune(r rune) rune {
	if isUpperRune(r) {
		return r + ('a' - 'A')
	}
	return r
}

func splitLetter(s string) []string {
	items := []rune(s)
	l := len(items)
	var results []string
	var item []rune
	for i := 0; i < l; i++ {
		if isLetterRune(items[i]) {
			item = append(item, items[i])
		} else {
			if len(item) > 0 {
				results = append(results, string(item))
				item = []rune{}
			}
		}
	}
	if len(item) > 0 {
		results = append(results, string(item))
	}
	return results
}
