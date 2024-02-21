package string

func isLetter[T byte | rune](val T) bool {
	return (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z')
}

func isUpper[T byte | rune](val T) bool {
	return val >= 'A' && val <= 'Z'
}

func isLower[T byte | rune](val T) bool {
	return val >= 'a' && val <= 'z'
}

func toUpper[T byte | rune](val T) T {
	if isLower(val) {
		return val - ('a' - 'A')
	}
	return val
}

func toLower[T byte | rune](val T) T {
	if isUpper(val) {
		return val + ('a' - 'A')
	}
	return val
}

func isLetterByte(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isUpperByte(r byte) bool {
	return r >= 'A' && r <= 'Z'
}

func isLowerByte(r byte) bool {
	return r >= 'a' && r <= 'z'
}

func toUpperByte(val byte) byte {
	if isLowerByte(val) {
		return val - ('a' - 'A')
	}
	return val
}

func toLowerByte(val byte) byte {
	if isUpperByte(val) {
		return val + ('a' - 'A')
	}
	return val
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
