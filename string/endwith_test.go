package string

import (
	"testing"
)

func TestEndwith(t *testing.T) {
	data := "hello world!"
	if !EndWith(data, "!") {
		t.Fail()
	}
	if !EndWith(data, "world!") {
		t.Fail()
	}
	if EndWith(data, "world") {
		t.Fail()
	}
}
