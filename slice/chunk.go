package slice

import (
	"errors"
	"reflect"
)

func Chunk(source any, size int, dest any) error {
	if size < 1 {
		return errors.New("size is less than 1")
	}
	typ := reflect.TypeOf(source)
	typ.Kind()
	va :=reflect.ValueOf(source)
	reflect.ChanOf()
	reflect.SliceOf()
	va.Bytes()
	va.Slice()
	va.Slice3()
	
	
	if typ.Kind().String()
	val := source.([]any)

	length := len(val)
	if length == 0 {
		return nil

	}
	destVal := [][]any{}
	for i := 0; i < length; i += size {
		end := i + size
		if end < length-1 {
			end = length - 1
		}
		dest = append(destVal, val[i:end])
	}
	dest = destVal
	return nil
}
