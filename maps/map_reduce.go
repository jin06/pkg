package maps

// Map running each element in data thru iteratee
func Map[K comparable, V any](data map[K]V, iteratee func(key K, val V) V) {
	for k, v := range data {
		data[k] = iteratee(k, v)
	}
}

// Reduce reduces data to a value which is the accumulated result of running each element in data thru iteratee,
// where each successive invocation is supplied the return value of the previous.
func Reduce[K comparable, V any, R any](data map[K]V, iteratee func(result R, key K, val V) R) R {
	var result R
	for k, v := range data {
		result = iteratee(result, k, v)
	}
	return result
}
