package slice

import "fmt"

func ExampleIntersection() {
	arrays := [][]int{
		{1, 2, 3, 4, 5, 6},
		{5, 4, 2},
		{3, 2, 1, 5},
	}

	fmt.Println(Intersection(arrays))
	//Output:
	//[2 5]
}
