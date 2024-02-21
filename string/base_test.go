package string

import "testing"

func TestIsLetter(t *testing.T) {
	letters := []byte{'a', 'g', 'z', 'A', 'B', 'Z'}
	nonLetters := []byte{'*', '!', ' ', '1', '0', '2'}
	for _, val := range letters {
		if !isLetter(val) || !isLetterByte(val) {
			t.Fail()
		}
	}
	for _, val := range nonLetters {
		if isLetter(val) || isLetterByte(val) {
			t.Fail()
		}
	}

	ch := "中文"
	for _, val := range ch {
		if isLetter(val) || isLetterRune(val) {
			t.Fail()
		}
	}

	mixed := "123abc中文"
	for i, val := range mixed {
		if i < 3 {
			if isLetter(val) || isLetterRune(val) {
				t.Fail()
			}
			continue
		}
		if i < 6 {
			if !isLetter(val) || !isLetterRune(val) {
				t.Fail()
			}
			continue
		}
		if i < 9 {
			if isLetter(val) || isLetterRune(val) {
				t.Fail()
			}
			continue
		}
	}
}

func TestIsUpperOrLower(t *testing.T) {
	var j byte = 'b'
	if isUpper(j) || !isLower(j) {
		t.Fail()
	}
	if isUpperByte(j) || !isLowerByte(j) {
		t.Fail()
	}
	var k byte = 'A'
	if !isUpper(k) || isLower(k) {
		t.Fail()
	}
	var l byte = '&'
	if isUpper(l) || isLower(l) {
		t.Fail()
	}
	var r rune = 'r'
	if isUpperRune(r) || !isLowerRune(r) {
		t.Fail()
	}
}

func TestToUpperOrLower(t *testing.T) {
	if toLower('!') != '!' {
		t.Fail()
	}
	if toLower('A') != 'a' {
		t.Fail()
	}
	if toLowerRune('B') != 'b' {
		t.Fail()
	}
	if toLowerByte('B') != 'b' {
		t.Fail()
	}
	if toUpper('a') != 'A' {
		t.Fail()
	}
	if toUpperRune('a') != 'A' {
		t.Fail()
	}
	if toUpperByte('a') != 'A' {
		t.Fail()
	}
}

var toLowerData = []byte{'a', 'B', '1'}

func BenchmarkToLower(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, val := range toLowerData {
			toLower(val)
		}
	}
}

func BenchmarkToLowerByte(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, val := range toLowerData {
			toLowerByte(val)
		}
	}
}

func BenchmarkToLowerRune(b *testing.B) {
	data := []rune{}
	for _, val := range toLowerData {
		data = append(data, rune(val))
	}
	for n := 0; n < b.N; n++ {
		for _, val := range data {
			toLowerRune(val)
		}
	}
}
