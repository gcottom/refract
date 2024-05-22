package refract

import (
	"fmt"
	"reflect"
)

func Append(slice any, elems ...any) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Ptr {
		val = reflect.ValueOf(slice).Elem()
	}
	if val.Kind() != reflect.Slice {
		return slice, fmt.Errorf("slice argument was not a slice")
	}
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

	newSlice := reflect.MakeSlice(reflect.SliceOf(val.Type().Elem()), val.Len()+elemsCount, val.Cap()+elemsCount)
	reflect.Copy(newSlice.Slice(0, val.Len()), val)

	counter := 0
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				newSlice.Index(val.Len() + counter).Set(reflect.Indirect(reflect.ValueOf(sliceItem)))
				counter++
			})
		} else {
			newSlice.Index(val.Len() + counter).Set(reflect.Indirect(reflect.ValueOf(elem)))
			counter++
		}

	}
	return newSlice.Interface(), nil
}

func Cap(v any) (int, error) {
	val := reflect.ValueOf(v)
	kind := val.Kind()
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Slice {
		return val.Cap(), nil
	}
	return 0, fmt.Errorf("kind: %v not supported by cap function", kind.String())
}

func Len(v any) (int, error) {
	val := reflect.ValueOf(v)
	kind := val.Kind()
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Slice || kind == reflect.Map || kind == reflect.String {
		return val.Len(), nil
	}
	return 0, fmt.Errorf("kind: %v not supported by len function", kind.String())
}

func Prepend(slice any, elems ...any) (any, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Ptr {
		val = reflect.ValueOf(slice).Elem()
	}
	if val.Kind() != reflect.Slice {
		return slice, fmt.Errorf("slice argument was not a slice")
	}
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
	newSlice := reflect.MakeSlice(reflect.SliceOf(val.Type().Elem()), val.Len()+elemsCount, val.Cap()+elemsCount)

	counter := 0
	for _, elem := range elems {
		if reflect.ValueOf(elem).Kind() == reflect.Slice {
			RangeOverSlice(elem, func(index int, sliceItem any) {
				newSlice.Index(counter).Set(reflect.Indirect(reflect.ValueOf(sliceItem)))
				counter++
			})
		} else {
			newSlice.Index(counter).Set(reflect.Indirect(reflect.ValueOf(elem)))
			counter++
		}
	}

	reflect.Copy(newSlice.Slice(elemsCount, val.Len()+elemsCount), val)
	return newSlice.Interface(), nil
}
