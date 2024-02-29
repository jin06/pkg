package slice

// IndexOf
// Gets the index at which the first occurrence of value is found in array. If from is negative, it's used as the offset from the end of array.
func IndexOf[T comparable](array []T, value T, from int) int {
	var start int
	if from >= 0 {
		start = from
	} else {
		start = len(array) + from
	}
	if start < 0 {
		return -1
	}
	for i := start; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}
	return -1
}
