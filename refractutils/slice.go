package refractutils

import (
	"errors"
	"fmt"

	"github.com/gcottom/refract/safereflect"
)

// GetSliceIndex takes a slice of any type and an index int. For dynamic types created with refract,
// it returns a modifiable (ptr) to the item at the index, otherwise returns the value at the index.
// This function returns an error if the index is out of bounds or the argument is not a slice.
func GetSliceIndex(slice any, index int) (any, error) {
	val := safereflect.ValueOf(slice)
	if val.Kind() == safereflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return nil, err
		}
		if index > length-1 {
			return nil, fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		vi, err := val.Index(index)
		if err != nil {
			return nil, err
		}
		return vi.Interface()
	}
	return nil, errors.New("slice argument was not a slice")
}

// GetSliceIndexValue takes a slice of any type and an index int. For dynamic types created with refract,
// it returns a copy of the item at the index (this copy will not be modifiable), otherwise returns the
// value at the index. This function returns an error if the index is out of bounds or the argument is not a slice.
func GetSliceIndexValue(slice any, index int) (any, error) {
	val := safereflect.ValueOf(slice)
	if val.Kind() == safereflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return nil, err
		}
		if index > length-1 {
			return nil, fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		vi, err := val.Index(index)
		if err != nil {
			return nil, err
		}
		vptr, err := vi.Interface()
		if err != nil {
			return nil, err
		}
		if safereflect.ValueOf(vptr).Kind() == safereflect.Pointer {
			vptre, err := safereflect.ValueOf(vptr).Elem()
			if err != nil {
				return nil, err
			}
			return vptre.Interface()
		}
		return vptr, nil
	}
	return nil, errors.New("slice argument was not a slice")
}

// SetSliceIndex takes a slice of any, the new value, and the index to put the new value at. This function returns
// an error if the index is out of bounds or the slice argument is not a slice
func SetSliceIndex(slice any, newValue any, index int) error {
	val := safereflect.ValueOf(slice)
	if val.Kind() == safereflect.Slice {
		length, err := Len(slice)
		if err != nil {
			return err
		}
		if index > length-1 {
			return fmt.Errorf("index: %d is out of range for slice with length: %d", index, length)
		}
		vi, err := val.Index(index)
		if err != nil {
			return err
		}
		if !vi.CanSet() {
			return fmt.Errorf("value at slice index: %d can not be set", index)
		}
		nval := safereflect.ValueOf(&newValue)
		if nval.Kind() == safereflect.Pointer {
			nval, err = nval.Elem()
			if err != nil {
				return err
			}
		}
		vi2, err := val.Index(index)
		if err != nil {
			return err
		}
		return vi2.Set(nval)
	}
	return errors.New("slice argument was not a slice")
}
