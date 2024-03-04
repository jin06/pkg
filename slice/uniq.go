package slice

// Uniq Creates a duplicate-free version of an array,in which only the first occurrence of each element is kept. The order of result values is determined by the order they occur in the array.
func Uniq[T comparable](array []T) (result []T) {
	result = make([]T, 0)
	cache := map[T]T{}
	for i := 0; i < len(array); i++ {
		if _, ok := cache[array[i]]; !ok {
			result = append(result, array[i])
			cache[array[i]] = array[i]
		}
	}
	return
}
