package maps

import (
	"fmt"
	"time"
)

func ExampleFilter() {
	first := time.Unix(1709531226, 0)
	second := first.Add(time.Hour)
	data := map[string]time.Time{
		"one":   first,
		"tow":   first,
		"later": second,
	}
	eqfirst := time.Unix(1709531226, 0)
	Filter(data, eqfirst)
	fmt.Println(data)
	//Output:
	//map[later:2024-03-04 14:47:06 +0800 CST]
}
