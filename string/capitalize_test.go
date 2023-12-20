package string

import "testing"

func TestCapitalize(t *testing.T) {
	s := "one OF a kid"
	d := "One of a Kid"
	if res := CapitalizeTitle(s); res != d {
		t.Log(res)
		t.Fail()
	}
}
