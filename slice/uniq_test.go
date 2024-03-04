package slice

import "fmt"

func ExampleUniq() {
	array := []string{"apple", "banana", "orange", "banana", "apple"}

	fmt.Println(Uniq(array))
	//Output:
	//[apple banana orange]
}
