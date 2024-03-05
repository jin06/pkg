package slice

import "fmt"

func ExampleMap() {
	data := []int{1, 2, 3, 4}

	result := Map(data, func(val int) int {
		result := val + 1
		return result
	})
	fmt.Println(result)
	//Output:
	//[2 3 4 5]
}

func ExampleReduce() {
	data := []int{1, 2, 3, 4}
	result := Reduce(data, func(result int, val int) int {
		return result + val
	})
	fmt.Println(result)
	//Output:
	//10
}
