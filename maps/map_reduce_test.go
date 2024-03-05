package maps

import "fmt"

func ExampleMap() {
	data := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	Map(data, func(key string, val int) int {
		result := val + 1
		return result
	})
	fmt.Println("one:", data["one"])
	fmt.Println("two:", data["two"])
	fmt.Println("three:", data["three"])
	fmt.Println("four:", data["four"])
	//Output:
	//one: 2
	//two: 3
	//three: 4
	//four: 5
}

func ExampleReduce() {
	data2 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	res2 := Reduce(data2, func(result int, key string, val int) int {
		return result + val
	})
	fmt.Println(res2)
	//Output:
	//10
}
