package slice

import (
	"math/rand"
	"time"
)

// Shuffle Returns the new shuffled slice.
func Shuffle[T any](array []T) (result []T) {
	result = make([]T, len(array))
	copy(result, array)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	l := len(array)
	for i := 0; i < l; i++ {
		j := r.Intn(l)
		if i != j {
			result[i], result[j] = result[j], result[i]
		}
	}
	return
}
