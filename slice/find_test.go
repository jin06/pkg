package slice

import (
	"fmt"
	"strings"
)

func ExampleFindIndex() {
	type user struct {
		Name string
		Age  int
	}
	datas := []user{
		{
			Name: "A",
			Age:  19,
		},
		{
			Name: "B",
			Age:  21,
		},
		{
			Name: "C",
			Age:  29,
		},
	}

	fmt.Println(FindIndex(datas, func(element user) bool {
		if element.Age > 20 && element.Age < 30 && element.Name != "" {
			return true
		}
		return false
	}))
	//Output:
	//1
}

func ExampleFindLastIndex() {
	datas := []string{
		"apple",
		"banana",
		"abc",
	}
	fmt.Println(FindLastIndex(datas, func(element string) bool {
		return strings.HasPrefix(element, "a")
	}))
	//Output:
	//2
}
