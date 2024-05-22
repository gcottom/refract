package refract

import (
	"fmt"
	"reflect"
)

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
	return nil, fmt.Errorf("slice argument was not a slice")
}

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
		return reflect.ValueOf(vptr).Elem().Interface(), nil
	}
	return nil, fmt.Errorf("slice argument was not a slice")
}

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
	return fmt.Errorf("slice argument was not a slice")
}
