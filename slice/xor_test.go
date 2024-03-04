package slice

import "fmt"

func ExampleXor() {
	arr1 := []string{"one", "two", ",", ".", "!", "ok"}
	arr2 := []string{"two", ".", "one", "hello", "123"}

	arr3 := []string{"one", ".", ",", "!", "kk", "09"}

	fmt.Println(Xor(arr1))
	fmt.Println(Xor(arr1, arr2))
	fmt.Println(Xor(arr1, arr2, arr3))
	//Output:
	//[]
	//[, ! ok]
	//[ok]
}
