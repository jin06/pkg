package slice

// Fill
// Fills elements of slice with value.
func Fill[T any](array []T, value T) (result []T) {
	result = make([]T, len(array))
	for i := 0; i < len(array); i++ {
		result[i] = value
	}
	return
}
