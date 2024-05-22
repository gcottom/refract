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
		var fields []reflect.StructField

		for i := 0; i < val.Index(index).Type().NumField(); i++ {
			fields = append(fields, reflect.StructField{Name: val.Index(index).Type().Field(i).Name, Type: val.Index(index).Field(i).Type(), Tag: val.Index(index).Type().Field(i).Tag})
		}
		structDef := NewStructDefinition(fields...)
		q := NewStructInstance(structDef)
		for i := 0; i < val.Index(index).Type().NumField(); i++ {
			err := SetStructFieldValue(q, val.Index(index).Type().Field(i).Name, val.Index(index).Field(i).Interface())
			if err != nil {
				return nil, err
			}
		}
		if err = SetSliceIndex(slice, q, index); err != nil {
			return nil, err
		}
		return q, nil
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
		nval := reflect.ValueOf(newValue)
		if nval.Kind() == reflect.Ptr {
			nval = nval.Elem()
		}
		val.Index(index).Set(nval)
		return nil
	}
	return fmt.Errorf("slice argument was not a slice")
}
