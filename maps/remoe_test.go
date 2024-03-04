package maps

import "fmt"

func ExampleRemove() {
	data := map[string]any{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	Remove(data, "three", "one")
	fmt.Println(data)
	//Output
	//map[one:1]
}
