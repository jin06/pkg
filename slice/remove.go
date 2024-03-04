package slice

// Remove removes all given values from array.
func Remove[T comparable](array []T, values []T) (result []T) {
	result = make([]T, 0)
	for i := 0; i < len(array); i++ {
		var have bool
		for j := 0; j < len(values); j++ {
			if values[j] == array[i] {
				have = true
				continue
			}
		}
		if !have {
			result = append(result, array[i])
		}
	}
	return result
}

// RemoveWhile is equivalent to DropWhile
func RemoveWhile[T any](array []T, fn func(e T) bool) (result []T) {
	return DropWhile[T](array, fn)
}
