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
