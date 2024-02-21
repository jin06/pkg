package slice

import (
	"fmt"
)

func ExampleChunk() {
	anySource := []any{1, 2, 3, "jk", true, 0.069}
	intSource := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	stringSource := []string{"a", "b", "c"}
	fmt.Println("any: ", Chunk(anySource, 2))
	fmt.Println("int:", Chunk(intSource, 5))
	fmt.Println("string:", Chunk(stringSource, 2))
}
