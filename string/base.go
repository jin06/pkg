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
