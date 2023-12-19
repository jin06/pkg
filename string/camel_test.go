package string

import "testing"

func TestCamelCase(t *testing.T) {
	if res := CamelCase("Foo Bar"); res != "fooBar" {
		t.Log(res)
		t.Fail()
	}
	if res := CamelCase("--foo-bar--"); res != "fooBar" {
		t.Log(res)
		t.Fail()
	}
	if res := CamelCase("__FOO_BAR__"); res != "fooBar" {
		t.Log(res)
		t.Fail()
	}
	if res := CamelCase("-Foo-BaR-"); res != "fooBar" {
		t.Log(res)
		t.Fail()
	}
	if res := CamelCase(""); res != "" {
		t.Log(res)
		t.Fail()
	}
	if res := CamelCase("898998"); res != "" {
		t.Log(res)
		t.Fail()
	}
}

func TestCamelCaseUpper(t *testing.T) {
	if res := CamelCaseUpper("-foo-bar-"); res != "FooBar" {
		t.Log(res)
		t.Fail()
	}
}
