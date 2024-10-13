package refractutils

import "github.com/gcottom/refract/safereflect"

func RangeOverSlice(slice any, f func(index int, sliceItem any)) error {
	length, err := Len(slice)
	if err != nil {
		return err
	}
	for i := 0; i < length; i++ {
		s, err := safereflect.ValueOf(slice).Index(i)
		if err != nil {
			return err
		}
		si, err := s.Interface()
		if err != nil {
			return err
		}
		f(i, si)
	}
	return nil
}

func RangeOverSliceReverse(slice any, f func(index int, sliceItem any)) error {
	length, err := Len(slice)
	if err != nil {
		return err
	}
	for i := length - 1; i >= 0; i-- {
		s, err := safereflect.ValueOf(slice).Index(i)
		if err != nil {
			return err
		}
		si, err := s.Interface()
		if err != nil {
			return err
		}
		f(i, si)
	}
	return nil
}

func RangeOverMap(m any, f func(counter int, key any, value any)) error {
	val := safereflect.ValueOf(m)
	mk, err := val.MapKeys()
	if err != nil {
		return err
	}
	for counter, k := range mk {
		key, err := k.Interface()
		if err != nil {
			return err
		}
		vmi, err := val.MapIndex(k)
		if err != nil {
			return err
		}
		value, err := vmi.Interface()
		if err != nil {
			return err
		}
		f(counter, key, value)
	}
	return nil
}
