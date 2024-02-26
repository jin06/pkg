package slice

import (
	"fmt"
	"testing"
)

func TestDifference(t *testing.T) {
	func() {
		first := []int{1, 2, 3, 4, 5, 6}
		second := []int{3, 5, 6}
		result := Difference(first, second)
		expect := []int{1, 2, 4}

		if len(result) == len(expect) {
			for i := 0; i < len(result); i++ {
				if !Equal(result[i], expect[i]) {
					t.Fail()
				}
			}
		} else {
			t.Fail()
		}
	}()
	func() {
		first := []string{"abc", "!!", "1900", "", "	"}
		second := []string{"abc", "1900"}
		result := Difference(first, second)
		expect := []string{"!!", "", "	"}

		if len(result) == len(expect) {
			for i := 0; i < len(result); i++ {
				if !Equal(result[i], expect[i]) {
					t.Fail()
				}
			}
		} else {
			t.Fail()
		}
	}()
}

func ExampleDifference() {
	first := []int{1, 2, 3}
	second := []int{2, 1, 4}
	fmt.Println(Difference(first, second))
	fmt.Println(Difference([]string{"19", "teck", "!"}, []string{"teck", "!"}))
	// Output:
	//[3]
	//[19]
}
