package string

func isLetter(r rune) bool {
	return (r >= 97 && r <= 122) || (r >= 65 && r <= 90)
}
