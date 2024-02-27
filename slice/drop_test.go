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

func ExampleDropWhile() {
	fn := func(i int) bool {
		return i%2 == 1
	}
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(DropWhile(array, fn))
	type vehicle struct {
		Brand string
		Class string
	}
	fn2 := func(v vehicle) bool {
		return v.Class == "suv" || v.Brand == "Ford"
	}
	array2 := []vehicle{
		{
			Brand: "Ford",
			Class: "car",
		},
		{
			Brand: "Ford",
			Class: "suv",
		},
		{
			Brand: "BMW",
			Class: "suv",
		},
		{
			Brand: "BMW",
			Class: "car",
		},
		{
			Brand: "Toyota",
			Class: "car",
		},
		{
			Brand: "Toyota",
			Class: "mpv",
		},
	}
	fmt.Println(DropWhile(array2, fn2))
	//Output:
	//[2 4 6 8]
	//[{BMW car} {Toyota car} {Toyota mpv}]
}
