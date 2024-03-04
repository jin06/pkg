package slice

// Reverse array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](array []T) (result []T) {
	count := len(array)
	result = make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = array[count-1-i]
	}
	return result
}
