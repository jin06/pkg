package slice

import "fmt"

func ExampleRemove() {
	array := []int{1, 3, 4, 5, 6, 3, 4}
	values := []int{1, 5, 6}
	fmt.Println(Remove(array, values...))
	//Output:
	//[3 4 3 4]
}
