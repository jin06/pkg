package slice

import (
	"math/rand"
	"time"
)

// Sample Gets n random elements at array.
func Sample[T any](array []T, n int) (result []T) {
	if n <= 0 || len(array) == 0 {
		return
	}
	result = make([]T, 0, n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	l := len(array)

	for i := 0; i < n; i++ {
		index := r.Intn(l)
		result = append(result, array[index])
	}
	return
}
