package slice

// Difference
// Creates an array of array values not included in the other given arrays using SameValueZero for equality comparisons.
// The order and references of result values are determined by the first array.
func Difference[T comparable](first []T, second []T) (result []T) {
	result = []T{}
	for i := 0; i < len(first); i++ {
		var found bool
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				found = true
				break
			}
		}
		if !found {
			result = append(result, first[i])
		}
	}
	return
}
