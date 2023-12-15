package string

import "testing"

func TestIsLetter(t *testing.T) {
	var a, g, z, A, B, Z, other byte = 'a', 'g', 'z', 'A', 'B', 'Z', '*'
	ch := "中文"
	if !isLetter(a) {
		t.Fail()
	}
	if !isLetter(g) {
		t.Fail()
	}
	if !isLetter(z) {
		t.Fail()
	}
	if !isLetter(A) {
		t.Fail()
	}
	if !isLetter(B) {
		t.Fail()
	}
	if !isLetter(Z) {
		t.Fail()
	}
	if isLetter(other) {
		t.Fail()
	}
	for i := 0; i < len(ch); i++ {
		if isLetter(ch[i]) {
			t.Fail()
		}
	}
}

func TestIsUpper(t *testing.T) {
	var j byte = 'b'
	var k byte = 'A'
	var l byte = '&'
	if isUpper(j) {
		t.Fail()
	}
	if !isUpper(k) {
		t.Fail()
	}
	if isUpper(l) {
		t.Fail()
	}
}
