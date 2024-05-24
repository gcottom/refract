package refract

import (
	"errors"
	"fmt"
	"reflect"
)

// Append is like the native go function append(). The refract.Append function takes a slice of any and variadic elems.
// It can be used with generic and dynamic slices as well as native types. Returns an error if slice is not a slice.
func Append(slice any, elems ...any) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Ptr {
		// if slice is a pointer to a slice, we get the concrete value
		val = reflect.ValueOf(slice).Elem()
	}
	if val.Kind() != reflect.Slice {
		// if slice is not a slice, return an error
		return slice, errors.New("slice argument was not a slice")
	}
	// get the reflect type of a pointer to the slice
	vtype := reflect.ValueOf(&slice)
	elemsCount := 0

	// count the elems in the argument
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				elemsCount++
			})
		} else {
			elemsCount++
		}
	}

	// make a new slice, (pointer to the slice type, so that its values can be edited later), with space for the new elems
	newSlice := reflect.MakeSlice(reflect.SliceOf(vtype.Type().Elem()), val.Len()+elemsCount, val.Cap()+elemsCount)

	// put all of the old elems in the new slice first
	if err := RangeOverSlice(slice, func(index int, sliceItem any) {
		v, err := GetSliceIndex(slice, index)
		if err != nil {
			return
		}
		newSlice.Index(index).Set(reflect.ValueOf(v))
	}); err != nil {
		return nil, err
	}

	// put all of the new elems in the new slice after the old elems
	counter := 0
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				newSlice.Index(val.Len() + counter).Set(reflect.ValueOf(sliceItem))
				counter++
			})
		} else {
			newSlice.Index(val.Len() + counter).Set(reflect.ValueOf(elem))
			counter++
		}

	}

	// return the newslice as interface so that we get the concrete value
	return newSlice.Interface(), nil
}

// Cap is like the native cap() function. It accepts an argument of any v, so it works with generic and dynamic types.
// If v is not an array, chan, slice, or a generic or dynamic type with one of these underylying types Cap returns an error.
func Cap(v any) (int, error) {
	val := reflect.ValueOf(v)
	kind := val.Kind()
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Slice {
		// if the kind of v is not array, chan, or slice, return an error, otherwise return the cap
		return val.Cap(), nil
	}
	return 0, fmt.Errorf("kind: %v not supported by cap function", kind.String())
}

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

// Prepend is the opposite of Append. Instead of adding elements at the end of the slice, it adds them at the beginning, in the order they appear in the variatic.
// Prepend accepts a slice any and a variadic elems, so it works with generic and dynamic types. If slice is not a slice, Prepend returns an error.
func Prepend(slice any, elems ...any) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Ptr {
		// if slice is a pointer we call elem to get the concrete value
		val = reflect.ValueOf(slice).Elem()
	}
	if val.Kind() != reflect.Slice {
		// if slice is not a slice return error
		return slice, errors.New("slice argument was not a slice")
	}

	// count all of the elems in the variadic
	elemsCount := 0
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				elemsCount++
			})
		} else {
			elemsCount++
		}
	}
	// get a value of the pointer to the slice
	vtype := reflect.ValueOf(&slice)
	// make a new slice, (pointer to the slice type, so that it can be edited later), that has enough space for the new elements
	newSlice := reflect.MakeSlice(reflect.SliceOf(vtype.Type().Elem()), val.Len()+elemsCount, val.Cap()+elemsCount)

	// add all of the new elems to the new slice
	counter := 0
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				newSlice.Index(counter).Set(reflect.ValueOf(sliceItem))
				counter++
			})
		} else {
			newSlice.Index(counter).Set(reflect.ValueOf(elem))
			counter++
		}
	}

	// add all of the old elems to the end of the new slice
	RangeOverSlice(slice, func(index int, sliceItem any) {
		v, err := GetSliceIndex(slice, index)
		if err != nil {
			return
		}
		newSlice.Index(elemsCount + index).Set(reflect.ValueOf(v))
	})

	// return the interface value of the new slice, therefore returning its concrete value
	return newSlice.Interface(), nil
}
