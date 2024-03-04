package slice

import "fmt"

func ExampleReverse() {
	array := []any{1, 2, 3, 4, 5, 6, "one", "two"}
	fmt.Println(Reverse(array))
	//Output:
	//[two one 6 5 4 3 2 1]
}
