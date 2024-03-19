package slice

import (
	"fmt"
	"testing"
)

func ExampleSample() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sample(data, 2))
}

func FuzzSample(f *testing.F) {
	f.Fuzz(func(t *testing.T, array []byte, n int) {
		res := Sample(array, n)
		t.Log(n)
		t.Log(res)
		if len(array) == 0 {
			if len(res) != 0 {
				t.Error()
			}
		} else {
			if n <= 0 {
				if len(res) != 0 {
					t.Error()
				}
			} else {
				if len(res) != n {
					t.Error()
				}
			}
		}
	})
}
