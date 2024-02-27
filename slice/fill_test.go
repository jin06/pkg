package slice

import (
	"fmt"
)

func ExampleFill() {
	arr := make([]string, 10)
	fmt.Println(Fill(arr, "default"))
	//Output:
	//[default default default default default default default default default default]
}
