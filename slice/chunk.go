package slice

// Chunk Creates an array of elements split into groups the length of size.If array can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](source []T, size int) (dest [][]T) {
	dest = [][]T{}
	for i := 0; i < len(source); i = i + size {
		dest = append(dest, source[i:i+size])
	}
	return
}
