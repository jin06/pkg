package string

// EndWith
// Checks if string ends with the given target string.
func EndWith(s string, target string) bool {
	l := len(s)
	tl := len(target)
	if tl == 0 {
		return true
	}
	if l == 0 || l < tl {
		return false
	}
	if s[l-tl:] == target {
		return true
	}
	return false
}
