package slice

// Drop
// Creates a slice of array with n elements dropped from the beginning while n >=0 and the ending while n < 0.
func Drop[T any](array []T, n int) (result []T) {
	length := len(array)
	if n < 0 {
		n = length + n
	}
	if length <= n || n < 0 {
		return array
	}
	result = append(result, array[:n]...)
	result = append(result, array[n+1:length]...)
	return
}

// DropWhile
// Creates a slice of array excluding elements dropped.
// Drop the elements that iterateFn result is true and return the other elements as a new slice.
func DropWhile[T any](array []T, iterateFn func(element T) bool) (result []T) {
	result = []T{}
	for i := 0; i < len(array); i++ {
		if !iterateFn(array[i]) {
			result = append(result, array[i])
		}
	}
	return
}
