package slice

// Map running each element in data thru iteratee, and return a new arrays.
func Map[V any](data []V, iteratee func(val V) V) []V {
	result := make([]V, len(data))
	for k, v := range data {
		result[k] = iteratee(v)
	}
	return result
}

// Reduce
func Reduce[V any, R any](data []V, iteratee func(result R, val V) R) R {
	var result R
	for _, v := range data {
		result = iteratee(result, v)
	}
	return result
}
