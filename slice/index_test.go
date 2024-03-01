package slice

import (
	"fmt"
)

func ExampleIndexOf() {
	arr1 := []int{2, 5, 1, 5, 67, 5, 6}
	fmt.Println(IndexOf(arr1, 5, 0))
	fmt.Println(IndexOf(arr1, 5, 2))
	fmt.Println(IndexOf(arr1, 5, -3))

	arr2 := []string{"a", "b", "a", "c"}
	fmt.Println(IndexOf(arr2, "b", 0))
	//Output:
	//1
	//3
	//5
	//1
}

func ExampleLastIndexOf() {
	arr1 := []int{2, 3, 2, 3, 1}
	fmt.Println(LastIndexOf(arr1, 3, 0))
	fmt.Println(LastIndexOf(arr1, 3, 2))
	fmt.Println(LastIndexOf(arr1, 3, -3))
	//Output:
	//3
	//1
	//1
}
