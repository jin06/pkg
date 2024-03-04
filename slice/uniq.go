package slice

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
