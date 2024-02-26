package slice

import (
	"fmt"
	"testing"
)

func TestDrop(t *testing.T) {
	arr1 := []int{}
	if len(Drop(arr1, 1)) > 0 {
		t.Error("drop nil error")
	}
	arr2 := []int{1, 2, 3, 4, 5}
	result2 := Drop(arr2, 2)
	if result2[2] != 4 {
		t.Error("drop 2 error result: ", result2)
	}
	arr3 := []string{"foo", "bar", "end"}
	result3 := Drop(arr3, 0)
	if result3[1] != "end" {
		t.Error("drop 0 error result: ", result3)
	}
	result4 := Drop(arr3, 2)
	if result4[len(result4)-1] != "bar" {
		t.Error("drop end error result: ", result4)
	}

	result5 := Drop(arr3, -1)
	if result5[len(result5)-1] != "bar" {
		t.Error("drop -1 error result: ", result5)
	}
}

func ExampleDrop() {
	array := []int{1, 3, 4, 5}
	fmt.Println(Drop(array, 1))

	array2 := []string{"start", "foo", "bar", "end"}
	fmt.Println(Drop(array2, -1))
	// Output:
	// [1 4 5]
	// [start foo bar]
}
