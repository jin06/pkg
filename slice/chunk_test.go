package slice

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	// t.SkipNow()
	var err error
	// var source = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var source = []any{1, 2, 3}
	var dest [][]any
	var sourceStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var destStrings [][]string

	if dest = Chunk(source, 3); err != nil {
		if len(dest) != 3 {
			t.Fail()
		} else {
			if dest[1][0] != 4 {
				t.Fail()
			}
		}
		t.Log(dest)
	}
	if destStrings = Chunk(sourceStrings, 3); err != nil {
		if len(destStrings) != 3 {
			t.Fail()
		} else {
			if destStrings[1][0] != "four" {
				t.Fail()
			}
		}
		t.Log(destStrings)
	}
}

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
