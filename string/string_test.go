package string

import "testing"

func TestKebabCase(t *testing.T) {
	if res := KebabCase("--Foo-Bar--"); res != "foo-bar" {
		t.Log(res)
		t.Fail()
	}
}
