package slice

import "testing"

func FuzzShuffle(f *testing.F) {
	f.Fuzz(func(t *testing.T, array []byte) {
		res := Shuffle(array)
		if len(array) != len(res) {
			t.Error()
		}
	})
}
