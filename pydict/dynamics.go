package pydict

import (
	"errors"
	"fmt"
	"reflect"
)

// Len is like the native len() function. It accepts an argument of any v, so it works with generic and dynamic types.
// If v is not an array, chan, slice, map, string, or a generic or dynamic type with one of these underlying types Len returns an error.
func Len(v any) (int, error) {
	val := reflect.ValueOf(v)
	kind := val.Kind()
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Slice || kind == reflect.Map || kind == reflect.String {
		// if the kind of v is not array, chan, slice, map, or string, return an error, otherwise return the length
		return val.Len(), nil
	}
	return 0, fmt.Errorf("kind: %v not supported by len function", kind.String())
}

// GetSliceIndex takes a slice of any type and an index int. For dynamic types created with refract,
// it returns a modifiable (ptr) to the item at the index, otherwise returns the value at the index.
// This function returns an error if the index is out of bounds or the argument is not a slice.
func GetSliceIndex(slice any, index int) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return nil, err
		}
		if index > length-1 {
			return nil, fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		return val.Index(index).Interface(), nil
	}
	return nil, errors.New("slice argument was not a slice")
}
