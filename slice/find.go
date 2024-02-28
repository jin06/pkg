package slice

// FindIndex
// This method returns the index of the first element fn returns truthy for instead of the element itself.
func FindIndex[T any](source []T, fn func(element T) bool) int {
	for i := 0; i < len(source); i++ {
		if fn(source[i]) {
			return i
		}
	}
	return -1
}

// FindLastIndex
// This method is like FindIndex, it iterates over elements of a slice from right to left.
func FindLastIndex[T any](source []T, fn func(element T) bool) int {
	for i := len(source) - 1; -1 < i; i-- {
		if fn(source[i]) {
			return i
		}
	}
	return -1
}
