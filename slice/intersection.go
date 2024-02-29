package slice

// Intersection
// Creates an array of unique values that are included in all given arrays. The order and references of result values are determined by the first array.
func Intersection[T comparable](arrays [][]T) (result []T) {
	if len(arrays) < 2 {
		return
	}
	result = arrays[0]
	for i := 1; i < len(arrays); i++ {
		result = intersection(result, arrays[i])
		if len(result) == 0 {
			break
		}
	}
	return
}

func intersection[T comparable](arr1 []T, arr2 []T) (result []T) {
	result = []T{}
	for i := 0; i < len(arr1); i++ {
		if IndexOf(arr2, arr1[i], 0) >= 0 {
			result = append(result, arr1[i])
		}
	}
	return
}
