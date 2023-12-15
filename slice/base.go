package slice

import "reflect"

func isSlice(val any) bool {
	t := reflect.TypeOf(val)
	if t.Kind().String() == reflect.Slice.String() {
		return true
	}
	return false
}

func test(val any) {

}
