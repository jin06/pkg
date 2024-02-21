package slice

func Chunk[T any](source []T, size int) (dest [][]T) {
	dest = [][]T{}
	for i := 0; i < len(source); i = i + size {
		dest = append(dest, source[i:i+size])
	}
	return
}
