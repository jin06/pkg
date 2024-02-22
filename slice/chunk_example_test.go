package slice

import (
	"fmt"
)

func ExampleChunk() {
	anySource := []any{1, 2, 3, "jk", true, 0.069}
	intSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stringSource := []string{"a", "b", "c"}
	fmt.Println(Chunk(anySource, 2))
	fmt.Println(Chunk(intSource, 5))
	fmt.Println(Chunk(stringSource, 2))
	fmt.Println(Chunk([]struct{}{}, 1))
	// Output:
	//[[1 2] [3 jk] [true 0.069]]
	//[[1 2 3 4 5] [6 7 8 9]]
	//[[a b] [c]]
	//[]
}
