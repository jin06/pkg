package string

// cameCase implement converts string to camel case.
func camelCase(source string, upper bool) string {
	l := len(source)
	result := make([]byte, 0)
	var lastFlag bool
	for i := 0; i < l; i++ {
		thisFlag := isLetter(source[i])
		if thisFlag {
			if !lastFlag {
				result = append(result, toUpper(source[i]))
			} else {
				result = append(result, toLower(source[i]))
			}
		}
		lastFlag = thisFlag
	}
	if !upper && len(result) > 0 {
		result[0] = toLower(result[0])
	}
	return string(result)
}

// CamelCase converts string to camel case and first letter is lower
func CamelCase(source string) string {
	return camelCase(source, false)
}

// CamelCaseUpper similar to CamelCase, but first letter is upper
func CamelCaseUpper(source string) string {
	return camelCase(source, true)
}
