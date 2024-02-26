package slice

// Drop
// Creates a slice of array with n elements dropped from the beginning while n >=0 and the endding while n < 0.
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
