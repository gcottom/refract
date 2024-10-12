package refractutils

import (
	"errors"
	"fmt"
	"reflect"
)

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

// GetSliceIndexValue takes a slice of any type and an index int. For dynamic types created with refract,
// it returns a copy of the item at the index (this copy will not be modifiable), otherwise returns the
// value at the index. This function returns an error if the index is out of bounds or the argument is not a slice.
func GetSliceIndexValue(slice any, index int) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return nil, err
		}
		if index > length-1 {
			return nil, fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		vptr := val.Index(index).Interface()
		if reflect.ValueOf(vptr).Kind() == reflect.Ptr {
			return reflect.ValueOf(vptr).Elem().Interface(), nil
		}
		return vptr, nil
	}
	return nil, errors.New("slice argument was not a slice")
}

// SetSliceIndex takes a slice of any, the new value, and the index to put the new value at. This function returns
// an error if the index is out of bounds or the slice argument is not a slice
func SetSliceIndex(slice any, newValue any, index int) error {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return err
		}
		if index > length-1 {
			return fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		if !val.Index(index).CanSet() {
			return fmt.Errorf("value at slice index: %d can not be set", index)
		}
		nval := reflect.ValueOf(&newValue)
		if nval.Kind() == reflect.Ptr {
			nval = nval.Elem()
		}
		val.Index(index).Set(nval)
		return nil
	}
	return errors.New("slice argument was not a slice")
}
