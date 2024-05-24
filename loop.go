package refract

import "reflect"

func RangeOverSlice(slice any, f func(index int, sliceItem any)) error {
	length, err := Len(slice)
	if err != nil {
		return err
	}
	for i := 0; i < length; i++ {
		s := reflect.ValueOf(slice).Index(i).Interface()
		f(i, s)
	}
	return nil
}

func RangeOverSliceReverse(slice any, f func(index int, sliceItem any)) error {
	length, err := Len(slice)
	if err != nil {
		return err
	}
	for i := length - 1; i >= 0; i-- {
		s := reflect.ValueOf(slice).Index(i).Interface()
		f(i, s)
	}
	return nil
}

func RangeOverMap(m any, f func(counter int, key any, value any)) error {
	val := reflect.ValueOf(m)
	for counter, k := range val.MapKeys() {
		key := k.Interface()
		value := val.MapIndex(k).Interface()
		f(counter, key, value)
	}
	return nil
}
