package slice

import (
	"testing"
	"time"
)

func TestEq(t *testing.T) {
	if Equal(1, 2) {
		t.Fail()
	}
	if Equal(1.2, 2.3) {
		t.Fail()
	}
	if Equal('a', 'A') {
		t.Fail()
	}
	if !Equal('k', 'k') {
		t.Fail()
	}
	if !Equal("hi,foo,bar", "hi,foo,bar") {
		t.Fail()
	}
	if Equal("hello,apple", "hello,banana") {
		t.Fail()
	}
	if Equal(time.Now(), time.Now().Add(time.Second)) {
		t.Fail()
	}
}
