package slice

// Xor Creates an array of unique values. The order of result values is determined by the order they occur in the arrays.
func Xor[T comparable](arrays ...[]T) (result []T) {
	if len(arrays) < 2 {
		return
	}
	result = arrays[0]
	for i := 1; i < len(arrays); i++ {
		result = xor(result, arrays[i])
		if len(result) == 0 {
			break
		}
	}
	return
}

func xor[T comparable](arr1 []T, arr2 []T) (result []T) {
	result = make([]T, 0)
	for i := 0; i < len(arr1); i++ {
		if IndexOf(arr2, arr1[i], 0) < 0 {
			result = append(result, arr1[i])
		}
	}
	return
}
